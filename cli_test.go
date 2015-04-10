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

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<items><item valid=\"true\" arg=\"u-ml-0\" uid=\"ml-0\"><title>.u-ml-0 { margin-left: 0; }</title><subtitle>Paste class name: u-ml-0</subtitle><icon>icon.png</icon></item></items>"
	assert.Equal(t, outStream.String(), expected)
}

func TestCLI_Run_MinusArgs(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./awc --", " ")

	status := cli.Run(args)
	assert.Equal(t, status, ExitCodeOK)

	expected := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<items><item valid=\"true\" arg=\"u-ti--9999\" uid=\"ti--9999\"><title>.u-ti--9999 { text-indent: -9999px; }</title><subtitle>Paste class name: u-ti--9999</subtitle><icon>icon.png</icon></item><item valid=\"true\" arg=\"u-ord--1\" uid=\"ord--1\"><title>.u-ord--1 { order: -1; }</title><subtitle>Paste class name: u-ord--1</subtitle><icon>icon.png</icon></item></items>"
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
