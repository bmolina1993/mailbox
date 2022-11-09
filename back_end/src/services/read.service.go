package services

import (
	"os"
)

// Lista carpeta/archivo
// valida si es carpeta
func ReadDirFile(path string) (map[string]bool, error) {
	var errMsg error
	files, err := os.ReadDir(path)

	if err != nil {
		errMsg = err
	}
	// map con nombre archivo
	// y bool si es carpeta
	data := make(map[string]bool)

	for _, item := range files {
		data[item.Name()] = item.IsDir()
	}

	return data, errMsg
}

// lectura por archivo
func ReadFile(filename string) (string, error) {
	var errMsg error
	file, err := os.ReadFile(filename)
	if err != nil {
		errMsg = err
	}
	return string(file), errMsg
}
