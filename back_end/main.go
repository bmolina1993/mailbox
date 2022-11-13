package main

import (
	"fmt"

	service "github.com/bmolina1993/mailbox/src/services"
)

func main() {
	pathFile := "./data/"

	//obtiene lista de carpeta de usuarios
	dirFile, _ := service.ReadDirFile(pathFile)

	//recorre cada carpeta de usuario
	//y sus sub-carpetas para hacer post api de data
	err := service.ExtractAllFolders(pathFile, dirFile)
	if err != nil {
		fmt.Println("ExtractAllFolders:", err)
	}

}
