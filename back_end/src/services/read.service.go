// package main
// package services
package service

import (
	"log"
	"os"
)

// Lista carpeta/archivo
// valida si es carpeta
func ReadDirFile(path string) map[string]bool {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	// map con nombre archivo
	// y bool si es carpeta
	data := make(map[string]bool)

	for _, item := range files {
		data[item.Name()] = item.IsDir()
	}

	return data
}

// lectura por archivo
func ReadFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
