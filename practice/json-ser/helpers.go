package main

import (
	"os"
	"path/filepath"
)

func WriteToMemory(data interface{}) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := "data.json"
	filePath := filepath.Join(currentDir, fileName)

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	return nil

}
