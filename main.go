package main

import (
	"os"
	"fmt"
	FileVerify "github.com/Logan-Garrett/FileCombiner/FileFunctions/file_funcs.go"
	// "strings"
)

func main() {
	// args := os.Args[1:]

	// Variable Definitions
	var input string
	files := []string{}

	fmt.Println("Please Provide Files: ")

	for {
		if input == "exit" {
			break
		}
		fmt.Scanln(&input)
		if FileVerify.IsFile(input) == false {
			continue
		} else {
        	files = append(files, input)
		}
        fmt.Print(":> ")
	}

	fmt.Println(files)
}
