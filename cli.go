package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ruedap/go-alfred"
	"github.com/ruedap/go-emma"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	str, err := find(args[1:])
	if err != nil {
		fmt.Fprintf(c.outStream, errorXML(err))
		return ExitCodeError
	}

	var decls []emma.Decl
	err = json.Unmarshal([]byte(str), &decls)
	if err != nil {
		fmt.Fprintf(c.outStream, errorXML(err))
		return ExitCodeError
	}

	res := response(decls)

	xml, err := res.ToXML()
	if err != nil {
		fmt.Fprintf(c.outStream, errorXML(err))
		return ExitCodeError
	}

	fmt.Fprint(c.outStream, xml)
	return ExitCodeOK
}

func errorXML(err error) string {
	title := fmt.Sprintf("Error: %v", err.Error())
	subtitle := "Emma.css Workflow Error"
	return alfred.ErrorXML(title, subtitle, title)
}

func find(terms []string) (string, error) {
	e := emma.NewEmma()
	return e.Find(terms).ToJSON()
}

func response(decls []emma.Decl) *alfred.Response {
	res := alfred.NewResponse()
	var arg, title, subtitle string
	for _, d := range decls {
		arg = fmt.Sprintf("u-%v", d.Snippet)
		title = fmt.Sprintf(".u-%v { %v: %v; }", d.Snippet, d.Property, d.Value)
		subtitle = fmt.Sprintf("Paste class name: u-%v", d.Snippet)
		ri := alfred.ResponseItem{
			Valid:    true,
			Arg:      arg,
			UID:      d.Snippet,
			Title:    title,
			Subtitle: subtitle,
			Icon:     "icon.png",
		}
		res.AddItem(&ri)
	}

	return res
}
