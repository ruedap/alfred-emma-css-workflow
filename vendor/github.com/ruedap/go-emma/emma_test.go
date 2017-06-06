package emma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmma_NewEmma(t *testing.T) {
	actual := NewEmma()
	assert.Equal(t, 856, len(actual.decls))
	assert.Equal(t, 0, len(actual.result))
}

func TestEmma_Find(t *testing.T) {
	e := NewEmma()
	terms := []string{"font-smoothing"}
	actual := e.Find(terms).result
	expected := []Decl{
		{"wkfsm-a", "-webkit-font-smoothing", "antialiased"},
		{"wkfsm-sa", "-webkit-font-smoothing", "subpixel-antialiased"},
		{"wkfsm-n", "-webkit-font-smoothing", "none"},
		{"mzfsm-g", "-moz-osx-font-smoothing", "grayscale"},
		{"mzfsm-u", "-moz-osx-font-smoothing", "unset"},
		{"fsm-a", "-webkit-font-smoothing", "antialiased"},
		{"@include emma-fsm-a;", "-webkit-font-smoothing", "antialiased"},
		{"fsm-a", "-moz-osx-font-smoothing", "grayscale"},
		{"@include emma-fsm-a;", "-moz-osx-font-smoothing", "grayscale"},
	}
	assert.Equal(t, expected, actual)
}

func TestEmma_ToCSS(t *testing.T) {
	decls := []Decl{
		{"pos-s", "position", "static"},
		{"pos-a", "position", "absolute"},
	}
	e := NewEmma()
	e.result = decls
	actual := e.ToCSS()
	expected := ".u-pos-s { position: static; }\n.u-pos-a { position: absolute; }\n"
	assert.Equal(t, expected, actual)

	decls = []Decl{{"ff-t", "font-family", `"Times New Roman", Times, Baskerville, Georgia, serif`}}
	e.result = decls
	actual = e.ToCSS()
	expected = ".u-ff-t { font-family: \"Times New Roman\", Times, Baskerville, Georgia, serif; }\n"
	assert.Equal(t, expected, actual)
}

func TestEmma_ToJSON(t *testing.T) {
	decls := []Decl{
		{"pos-s", "position", "static"},
		{"pos-a", "position", "absolute"},
	}

	e := NewEmma()
	e.result = decls
	actual, err := e.ToJSON()
	assert.Nil(t, err)

	expected := "[{\"snippet\":\"pos-s\",\"property\":\"position\",\"value\":\"static\"},{\"snippet\":\"pos-a\",\"property\":\"position\",\"value\":\"absolute\"}]"
	assert.Equal(t, expected, actual)
}

func TestEmma_ToJSON_Zero(t *testing.T) {
	decls := []Decl{}

	e := NewEmma()
	e.result = decls
	actual, err := e.ToJSON()
	assert.Nil(t, err)

	expected := "[]"
	assert.Equal(t, expected, actual)
}

func TestEmma_parse(t *testing.T) {
	src := `
rules:
  props:
    -
      name: width
      abbr: w
      group: width
      values:
        -
          name: auto
          abbr: a
        -
          name: '0'
          abbr: '0'
        -
          name: 1px
          abbr: '1'
        -
          name: 1%
          abbr: 1p
        -
          name: 100%
          abbr: 100p
        -
          name: 25vw
          abbr: 25vw
  mixins:
    -
      name: clearfix
      abbr: cf
      desc: 'Clearfix (Contain floats)'
      group: float
      decls:
        -
          prop: ''
          value: '&::after { content: ""; display: table; clear: both; }'
    -
      name: text-truncation
      abbr: tetr
      desc: 'Text truncation (ellipsis)'
      group: text
      decls:
        -
          prop: max-width
          value: 100%
        -
          prop: overflow
          value: hidden
`

	e := NewEmma()
	e.setSrc(src)
	actual, err := e.parse()
	assert.Nil(t, err)

	expected := []Decl{
		{
			Snippet:  "w-a",
			Property: "width",
			Value:    "auto",
		},
		{
			Snippet:  "w0",
			Property: "width",
			Value:    "0",
		},
		{
			Snippet:  "w1",
			Property: "width",
			Value:    "1px",
		},
		{
			Snippet:  "w1p",
			Property: "width",
			Value:    "1%",
		},
		{
			Snippet:  "w100p",
			Property: "width",
			Value:    "100%",
		},
		{
			Snippet:  "w25vw",
			Property: "width",
			Value:    "25vw",
		},
		{
			Snippet:  "cf",
			Property: "",
			Value:    "&::after { content: \"\"; display: table; clear: both; }",
		},
		{
			Snippet:  "@include emma-cf;",
			Property: "",
			Value:    "&::after { content: \"\"; display: table; clear: both; }",
		},
		{
			Snippet:  "tetr",
			Property: "max-width",
			Value:    "100%",
		},
		{
			Snippet:  "@include emma-tetr;",
			Property: "max-width",
			Value:    "100%",
		},
		{
			Snippet:  "tetr",
			Property: "overflow",
			Value:    "hidden",
		},
		{
			Snippet:  "@include emma-tetr;",
			Property: "overflow",
			Value:    "hidden",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestEmma_parse_Blank(t *testing.T) {
	e := NewEmma()
	e.setSrc("")
	actual, err := e.parse()
	assert.NotNil(t, err)

	expected := []Decl{}
	assert.Equal(t, expected, actual)
}

func TestEmma_parseProps(t *testing.T) {
	props := []TEmmaDocProp{
		{
			Name:  "top",
			Abbr:  "t",
			Group: "display",
			Values: []TEmmaDocPropValue{
				{
					Name: "auto",
					Abbr: "a",
				},
				{
					Name: "0",
					Abbr: "0",
				},
			},
		},
	}

	actual, err := parseProps(props)
	assert.Nil(t, err)

	expected := []Decl{
		{
			Snippet:  "t-a",
			Property: "top",
			Value:    "auto",
		},
		{
			Snippet:  "t0",
			Property: "top",
			Value:    "0",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestEmma_parseMixins(t *testing.T) {
	mixins := []TEmmaDocMixin{
		{
			Name:  "text-hiding",
			Abbr:  "tehi",
			Desc:  "Text hiding",
			Group: "text",
			Decls: []TEmmaDocMixinDecl{
				{
					Prop:  "overflow",
					Value: "hidden",
				},
				{
					Prop:  "text-indent",
					Value: "200%",
				},
			},
		},
	}

	actual, err := parseMixins(mixins)
	assert.Nil(t, err)

	expected := []Decl{
		{
			Snippet:  "tehi",
			Property: "overflow",
			Value:    "hidden",
		},
		{
			Snippet:  "@include emma-tehi;",
			Property: "overflow",
			Value:    "hidden",
		},
		{
			Snippet:  "tehi",
			Property: "text-indent",
			Value:    "200%",
		},
		{
			Snippet:  "@include emma-tehi;",
			Property: "text-indent",
			Value:    "200%",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestEmma_contains_True(t *testing.T) {
	d := Decl{"pos-s", "position", "static"}
	actual := contains(d, []string{"s-s"})
	assert.True(t, actual)

	actual = contains(d, []string{"s", "s", "s"})
	assert.True(t, actual)

	actual = contains(d, []string{"static", "position", "pos-s"})
	assert.True(t, actual)
}

func TestEmma_contains_False(t *testing.T) {
	d := Decl{"pos-s", "position", "static"}
	actual := contains(d, []string{"pos-a"})
	assert.False(t, actual)

	actual = contains(d, []string{"s-s", "pos-a"})
	assert.False(t, actual)

	actual = contains(d, []string{"s", "s", "z"})
	assert.False(t, actual)
}

func TestEmma_containsDecl_True(t *testing.T) {
	d := Decl{"pos-s", "position", "static"}
	actual := containsDecl(d, "s-s")
	assert.True(t, actual)

	actual = containsDecl(d, "ti")
	assert.True(t, actual)

	actual = containsDecl(d, "")
	assert.True(t, actual)
}

func TestEmma_containsDecl_False(t *testing.T) {
	d := Decl{"pos-s", "position", "static"}
	actual := containsDecl(d, "pos-a")
	assert.False(t, actual)

	actual = containsDecl(d, "spo")
	assert.False(t, actual)

	actual = containsDecl(d, "ss")
	assert.False(t, actual)
}

func TestEmma_isNumeric_True(t *testing.T) {
	assert.True(t, isNumeric("0"))
	assert.True(t, isNumeric("10px"))
	assert.True(t, isNumeric("-3"))
}

func TestEmma_isNumeric_False(t *testing.T) {
	assert.False(t, isNumeric("a1"))
	assert.False(t, isNumeric("b-"))
	assert.False(t, isNumeric("-moz"))
}

func TestEmma_generateAbbr(t *testing.T) {
	assert.Equal(t, "w-a", generateAbbr("w", "a"))
	assert.Equal(t, "w0", generateAbbr("w", "0"))
	assert.Equal(t, "w1", generateAbbr("w", "1"))
	assert.Equal(t, "w25vw", generateAbbr("w", "25vw"))
	assert.Equal(t, "ti-9999", generateAbbr("ti", "-9999"))
}
