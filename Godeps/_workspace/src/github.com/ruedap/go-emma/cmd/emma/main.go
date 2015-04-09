package main

import (
	"fmt"
	"os"

	"github.com/ruedap/go-emma"
)

func main() {
	e := emma.NewEmma()
	str, err := e.Find(os.Args[1:]).ToJSON()
	if err != nil {
		fmt.Println("Failed to output json.")
		return
	}

	fmt.Println(str)
}
