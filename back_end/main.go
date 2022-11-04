package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

// lectura por archivo
func ReadFile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

// Lista carpeta/archivo
// valida si es carpeta
func ReadDirFile(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range files {
		fmt.Printf("IsDir: %t | Name: %s \n", item.IsDir(), item.Name())
	}
}

func main() {
	//pathFile := "./data"
	//fileName := pathFile + "/allen-p/_sent_mail/3."

	//prueba de funciones
	//ReadFilePerLine(fileName)
	//ReadFile(fileName)
	ReadDirFile("./data/shively-h/")

}
