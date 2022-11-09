package main

import (
	"fmt"

	service "github.com/bmolina1993/mailbox/src/services"
)

func main() {

	pathFile := "./data/"

	//obtiene lista de carpeta de usuarios
	dirFile, _ := service.ReadDirFile(pathFile)
	/**/
	fmt.Println("--> Carpetas de usuarios <--")
	fmt.Println(dirFile)

	//recorre cada carpeta de usuario
	//y sus sub-carpetas para hacer post api de data
	service.ExtractAllFolders(pathFile, dirFile)

	//borra todos los documentos[para pruebas][x]
	//recorre cara carpeta de usuario
	/*
		for userFolder := range dirFile {
			service.DeleteDocuments(userFolder)
		}

			//por carpeta de usuario especifico
			service.DeleteDocuments("stokley-c")
	*/
}
