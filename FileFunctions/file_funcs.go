package file_funcs

import (
	"os"
)

func IsFile(input string) bool {
	_, err := os.Stat(input)
	if os.IsNotExist(err) && input != "exit" {
        fmt.Println("File not Found !!")
        return false
    }

	return true
}
