package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// Get Home dir
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not get Home Dir Path")
		return
	}
	filePath := filepath.Join(home, "WorkSpace/Go/httpfromtcp/messages.txt")

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error while oping file")
		return
	}
	defer file.Close()

	data := make([]byte, 8)
	for true {
		_, err = file.Read(data)
		if err != nil {
			fmt.Printf("Error while Reading file")
			return
		}
		fmt.Printf("read: %s", data)
	}

}
