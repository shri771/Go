package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Students struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class string `json:"class"`
	Cgr   int    `json:"cgr"`
}

func main() {
	home, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get home dir: %v", err)
		return
	}

	fileName := "tst.json"
	filePath := filepath.Join(home, fileName)

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Colud not Open file: %v", err)
		return
	}

	students := Students{}
	err = json.NewDecoder(file).Decode(&students)
	if err != nil {
		fmt.Printf("Colud not Decode: %v", err)
		return
	}

}
