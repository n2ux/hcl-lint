package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl"
	"github.com/mitchellh/colorstring"
)

func main() {
	var parseErr int
	if len(os.Args) == 2 && os.Args[1] == "-" {
		bytes, err := ioutil.ReadAll(os.Stdin)
		_, err = hcl.Parse(string(bytes))
		if err != nil {
			colorstring.Printf("[red]Error parsing stdin: %s\n", err)
			parseErr = 1
		} else {
			colorstring.Printf("[green]OK!\n")
		}
	} else {
		for i, arg := range os.Args {
			if i == 0 {
				continue
			}
			search := arg
			if info, err := os.Stat(arg); err == nil && info.IsDir() {
				search = fmt.Sprintf("%s/*.tf", arg)
			}
			files, err := filepath.Glob(search)
			if err != nil {
				colorstring.Printf("[red]Error finding files: %s", err)
			}
			for _, filename := range files {
				fmt.Printf("Checking %s ... ", filename)
				file, err := ioutil.ReadFile(filename)
				if err != nil {
					colorstring.Printf("[red]Error reading file: %s\n", err)
					break
				}
				_, err = hcl.Parse(string(file))
				if err != nil {
					colorstring.Printf("[red]Error parsing file: %s\n", err)
					parseErr = 1
					break
				}
				colorstring.Printf("[green]OK!\n")
			}
		}
	}
	os.Exit(parseErr)
}
