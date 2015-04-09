package emma

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Emma struct {
	src    string
	decls  []Decl
	result []Decl
}

// Single declaration
type Decl struct {
	Snippet  string `json:"snippet"`
	Property string `json:"property"`
	Value    string `json:"value"`
}

func NewEmma() *Emma {
	e := new(Emma)
	e.setSrc(Src)
	e.result = []Decl{}

	return e
}

func (e *Emma) Find(terms []string) *Emma {
	var res []Decl
	for _, d := range e.decls {
		if contains(d, terms) {
			res = append(res, d)
		}
	}

	e.result = res
	return e
}

func (e *Emma) ToCSS() string {
	var str string
	for _, d := range e.result {
		str += fmt.Sprintf(".u-%s { %s: %s; }\n", d.Snippet, d.Property, d.Value)
	}

	return str
}

func (e *Emma) ToJSON() (string, error) {
	if len(e.result) == 0 {
		return "[]", nil
	}

	b, err := json.Marshal(e.result)
	return string(b), err
}

func (e *Emma) parse() ([]Decl, error) {
	re := regexp.MustCompile(`\s+\((.+?)\,(.+?)\,(.+)\)\,.*`)
	res := re.FindAllStringSubmatch(e.src, -1)
	var dec Decl
	var ret []Decl

	if len(res) < 1 {
		return []Decl{}, errors.New("failed to parse source file")
	}

	for _, sl := range res {
		if len(sl) != 4 {
			continue
		}

		s := strings.TrimSpace(sl[3])
		switch {
		case s[0] == `'`[0] && s[len(s)-1] == `'`[0]:
			s = strings.Trim(s, `'`)
		case s[0] == `"`[0] && s[len(s)-1] == `"`[0]:
			s = strings.Trim(s, `"`)
		}

		dec = Decl{
			Snippet:  strings.TrimSpace(sl[1]),
			Property: strings.TrimSpace(sl[2]),
			Value:    s,
		}
		ret = append(ret, dec)
	}

	return ret, nil
}

func (e *Emma) setSrc(src string) *Emma {
	e.src = src
	decls, _ := e.parse()
	e.decls = decls

	return e
}

func contains(d Decl, terms []string) bool {
	for _, t := range terms {
		if !containsDecl(d, t) {
			return false
		}
	}

	return true
}

func containsDecl(d Decl, term string) bool {
	if strings.Contains(d.Snippet, term) {
		return true
	}

	if strings.Contains(d.Property, term) {
		return true
	}

	if strings.Contains(d.Value, term) {
		return true
	}

	return false
}
