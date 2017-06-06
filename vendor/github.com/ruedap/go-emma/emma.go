package emma

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type TEmmaDocVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type TEmmaDocProp struct {
	Name   string `yaml:"name"`
	Abbr   string `yaml:"abbr"`
	Group  string `yaml:"group"`
	Values []TEmmaDocPropValue
}

type TEmmaDocPropValue struct {
	Name string `yaml:"name"`
	Abbr string `yaml:"abbr"`
}

type TEmmaDocMixin struct {
	Name  string `yaml:"name"`
	Abbr  string `yaml:"abbr"`
	Desc  string `yaml:"desc"`
	Group string `yaml:"group"`
	Decls []TEmmaDocMixinDecl
}

type TEmmaDocMixinDecl struct {
	Prop  string `yaml:"prop"`
	Value string `yaml:"value"`
}

type TEmmaDocRule struct {
	Props  []TEmmaDocProp
	Mixins []TEmmaDocMixin
}

type TEmmaDoc struct {
	Ver   string `yaml:"ver"`
	Vars  []TEmmaDocVar
	Rules TEmmaDocRule
}

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
	data := FSMustString(false, "/assets/emma-data.yml")
	e.setSrc(data)
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
	doc := TEmmaDoc{}
	err := yaml.Unmarshal([]byte(e.src), &doc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var resultProps []Decl
	resultProps, err = parseProps(doc.Rules.Props)
	if err != nil {
		return []Decl{}, errors.New("failed to parse source file")
	}

	var resultMixins []Decl
	resultMixins, err = parseMixins(doc.Rules.Mixins)
	if err != nil {
		return []Decl{}, errors.New("failed to parse source file")
	}

	return append(resultProps, resultMixins...), nil
}

func parseProps(props []TEmmaDocProp) ([]Decl, error) {
	var result []Decl

	if len(props) < 1 {
		return []Decl{}, errors.New("failed to parse source file")
	}

	for _, prop := range props {
		for _, value := range prop.Values {
			result = append(result, Decl{
				Snippet:  generateAbbr(prop.Abbr, value.Abbr),
				Property: prop.Name,
				Value:    value.Name,
			})
		}
	}

	return result, nil
}

func parseMixins(mixins []TEmmaDocMixin) ([]Decl, error) {
	var result []Decl

	if len(mixins) < 1 {
		return []Decl{}, errors.New("failed to parse source file")
	}

	for _, mixin := range mixins {
		for _, decl := range mixin.Decls {
			result = append(result, Decl{
				Snippet:  mixin.Abbr,
				Property: decl.Prop,
				Value:    decl.Value,
			})
			result = append(result, Decl{
				Snippet:  "@include emma-" + mixin.Abbr + ";",
				Property: decl.Prop,
				Value:    decl.Value,
			})
		}
	}

	return result, nil
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

func isNumeric(str string) bool {
	re := regexp.MustCompile(`^(\d|-\d)`)
	return re.MatchString(str)
}

func generateAbbr(propAbbr, valueAbbr string) string {
	if isNumeric(valueAbbr) {
		return propAbbr + valueAbbr
	}
	return propAbbr + "-" + valueAbbr
}
