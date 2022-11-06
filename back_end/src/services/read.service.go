package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

// lectura de archivo (linea a linea)
func ReadFilePerLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		//fmt.Println(err.Error())
		log.Fatal(err.Error())
	}

	//cierra archivo luego de usarlo
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), ":")
		fmt.Println(len(raw))
	}
}
