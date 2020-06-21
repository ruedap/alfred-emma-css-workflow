package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ruedap/go-alfred"
	"github.com/ruedap/go-emma"
	"github.com/stretchr/testify/assert"
)

func TestCLI_Run(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./awc margin-left 0", " ")

	status := cli.Run(args)
	assert.Equal(t, status, ExitCodeOK)

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<items><item valid=\"true\" arg=\"ml0\" uid=\"ml0\"><title>ml0 { margin-left: 0; }</title><subtitle>Paste class name: ml0</subtitle><icon>icon.png</icon></item><item valid=\"true\" arg=\"mx0\" uid=\"mx0\"><title>mx0 { margin-left: 0; }</title><subtitle>Paste class name: mx0</subtitle><icon>icon.png</icon></item><item valid=\"true\" arg=\"@include emma-mx0;\" uid=\"@include emma-mx0;\"><title>@include emma-mx0; { margin-left: 0; }</title><subtitle>Paste class name: @include emma-mx0;</subtitle><icon>icon.png</icon></item></items>"
	assert.Equal(t, outStream.String(), expected)
}

func TestCLI_Run_MinusArgs(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./awc -1", " ")

	status := cli.Run(args)
	assert.Equal(t, status, ExitCodeOK)

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<items><item valid=\"true\" arg=\"ord-1\" uid=\"ord-1\"><title>ord-1 { order: -1; }</title><subtitle>Paste class name: ord-1</subtitle><icon>icon.png</icon></item><item valid=\"true\" arg=\"fx0-1-a\" uid=\"fx0-1-a\"><title>fx0-1-a { flex: 0 1 auto; }</title><subtitle>Paste class name: fx0-1-a</subtitle><icon>icon.png</icon></item><item valid=\"true\" arg=\"fx1-1-a\" uid=\"fx1-1-a\"><title>fx1-1-a { flex: 1 1 auto; }</title><subtitle>Paste class name: fx1-1-a</subtitle><icon>icon.png</icon></item></items>"
	assert.Equal(t, outStream.String(), expected)
}

func TestCLI_Run_NoResult(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./awc emma.css", " ")

	status := cli.Run(args)
	assert.Equal(t, status, ExitCodeOK)

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<items></items>"
	assert.Equal(t, outStream.String(), expected)
}

func TestCLI_errorXML(t *testing.T) {
	err := errors.New("Foo")
	actual := errorXML(err)

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<items><item valid=\"false\" arg=\"Error: Foo\" uid=\"error\"><title>Error: Foo</title><subtitle>Emma.css Workflow Error</subtitle><icon>/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns</icon></item></items>"
	assert.Equal(t, actual, expected)
}

func TestCLI_find(t *testing.T) {
	actual, err := find([]string{"foo"})
	assert.Nil(t, err)

	expected := `[{"snippet":"d-tbfg","property":"display","value":"table-footer-group"}]`
	assert.Equal(t, actual, expected)
}

func TestCLI_response(t *testing.T) {
	actual := response([]emma.Decl{})
	expected := new(alfred.Response)
	expected.Items = []alfred.ResponseItem{}
	assert.Equal(t, actual, expected)
}
