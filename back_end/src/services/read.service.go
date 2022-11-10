package services

import (
	"os"
)

// Lista carpeta/archivo
// valida si es carpeta
func ReadDirFile(path string) (map[string]bool, error) {

	files, err := os.ReadDir(path)

	if err != nil {
		return map[string]bool{}, err
	}
	// map con nombre archivo
	// y bool si es carpeta
	data := make(map[string]bool)

	for _, item := range files {
		data[item.Name()] = item.IsDir()
	}

	return data, nil
}

// lectura por archivo
func ReadFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
