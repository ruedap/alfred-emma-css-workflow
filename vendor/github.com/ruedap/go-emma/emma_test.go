package emma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmma_NewEmma(t *testing.T) {
	actual := NewEmma()
	assert.Equal(t, 714, len(actual.decls))
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
      name: position
      abbr: pos
      group: display
      values:
        -
          name: static
          abbr: s
        -
          name: relative
          abbr: r
`

	e := NewEmma()
	e.setSrc(src)
	actual, err := e.parse()
	assert.Nil(t, err)

	expected := []Decl{
		{
			Snippet:  "pos-s",
			Property: "position",
			Value:    "static",
		},
		{
			Snippet:  "pos-r",
			Property: "position",
			Value:    "relative",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestEmma_parse_FontFamily(t *testing.T) {
	src := `
rules:
  props:
    -
      name: font-family
      abbr: ff
      group: text
      values:
        -
          name: '"Times New Roman", Times, Baskerville, Georgia, serif'
          abbr: t
`

	e := NewEmma()
	e.setSrc(src)
	actual, err := e.parse()
	assert.Nil(t, err)

	expected := []Decl{
		{
			Snippet:  "ff-t",
			Property: "font-family",
			Value:    `"Times New Roman", Times, Baskerville, Georgia, serif`,
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
