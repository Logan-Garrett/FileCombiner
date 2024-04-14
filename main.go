package main

import (
	"os"
	"fmt"
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
		_, err := os.Stat(input)
		if err != nil {
			if os.IsNotExist(err) && input != "exit" {
            	fmt.Println("File not Found !!")
        	}
		} else {
        	files = append(files, input)
		}
        fmt.Print(":> ")
	}

	fmt.Println(files)
}
