package file_funcs

import (
	"os"
	// "fmt"
)

func IsFile(input string) bool {
	_, err := os.Stat(input)
	if input == "exit" || os.IsNotExist(err) {
        return false
    } else {
    	return true
    }
}
