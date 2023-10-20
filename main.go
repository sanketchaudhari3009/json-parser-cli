package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sanketchaudhari3009/json-parser-go/parser"
)

func main() {
	filePath := flag.String("file", "", "The path of the JSON file to be validated")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide the path of the JSON file to be validated")
		os.Exit(1)
	}

	jsonBytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println("Could not read the file:", err)
		os.Exit(1)
	}

	err = parser.ParseJSON(string(jsonBytes))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("JSON is valid.")
	os.Exit(0)
}
