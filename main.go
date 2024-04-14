package main

import (
	// "os"
	"fmt"
	FileVerify "FileCombiner/file_funcs"
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
			fmt.Println("File not Found!")
			fmt.Print(":> ")
		} else {
        	files = append(files, input)
         	fmt.Print(":> ")
		}
	}

	fmt.Println(files)
}
