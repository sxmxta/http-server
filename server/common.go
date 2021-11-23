package server

import (
	"fmt"
	"os"
)

func PathExists(path string) bool {
	fmt.Println(path)
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}