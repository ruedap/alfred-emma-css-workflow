package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ruedap/go-alfred"
	"github.com/ruedap/go-emma"
)

func main() {
	e := emma.NewEmma()
	str, err := e.Find(os.Args[1:]).ToJSON()
	if err != nil {
		fmt.Println(err)
		return
	}

	var decls []emma.Decl
	err = json.Unmarshal([]byte(str), &decls)

	res := alfred.NewResponse()
	for _, d := range decls {
		ri := alfred.ResponseItem{
			Valid:    true,
			Arg:      fmt.Sprintf("u-%v", d.Snippet),
			UID:      d.Snippet,
			Title:    fmt.Sprintf(".u-%v { %v: %v; }", d.Snippet, d.Property, d.Value),
			Subtitle: fmt.Sprintf("Paste class name: u-%v", d.Snippet),
			Icon:     "icon.png",
		}
		res.AddItem(&ri)
	}

	xml, err := res.ToXML()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(xml)
}
