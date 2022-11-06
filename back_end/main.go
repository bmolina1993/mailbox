package main

import (
	"fmt"

	service "github.com/bmolina1993/mailbox/src/services"
)

func main() {
	pathFile := "./data/"

	/*
		fileName := pathFile + "allen-p/_sent_mail/9."
		//prueba de funciones
		//ReadFilePerLine(fileName)
		data := service.ReadFile(fileName)
		props, body := service.ExtractData(data)

		fmt.Println("--> props <--")
		fmt.Println(props)

		fmt.Println("--> body <--")
		fmt.Println(body)

				fmt.Println("--> prop - [To] <--")
				to := extractPropTo(props)
				fmt.Println(to)

				fmt.Println("--> prop - [Date] <--")
				date := extractPropDate(props)
				fmt.Println(date)

				fmt.Println("--> prop - [From] <--")
				from := extractPropFrom(props)
				fmt.Println(from)

			fmt.Println("--> prop - [Subject] <--")
			subject := extractPropSubject(props)
			fmt.Println(subject)
	*/

	fmt.Println("--> Carpetas <--")
	//caso de usuario con archivo suelto en carpeta usuario
	//dirFile := ReadDirFile("./data/shively-h/")
	dirFile := service.ReadDirFile(pathFile)
	fmt.Println(dirFile)

	// TEMPORAL [X]
	// filtra carpeta
	for fileToDelete := range dirFile {
		if fileToDelete != "motley-m" {
			delete(dirFile, fileToDelete)
		}
	}
	fmt.Println("\n--> filtra Carpeta <--")
	fmt.Println(dirFile)
	// TEMPORAL [X]

	/**/
	fmt.Println("\n--> Recorre carpeta por usuario <--")

	//accede a carpeta de usuarios
	for userFolder := range dirFile {
		// iteracion por usuario,
		// se omiten archivos por fuera de carpetas
		pathPerUser := pathFile + userFolder
		dirPerUser := service.ReadDirFile(pathPerUser)

		// TEMPORAL [X]
		// filtra carpeta
		for folderToDelete := range dirPerUser {
			if folderToDelete != "sent_items" {
				delete(dirPerUser, folderToDelete)
			}
		}
		fmt.Println("\n--> filtra Carpeta interna de usuario <--")
		fmt.Println(dirPerUser)
		// TEMPORAL [X]

		fmt.Println("\n--> recorre cada archivo por carpeta de usuario <--")
		for folderName, isDir := range dirPerUser {
			if isDir {
				//fmt.Printf("folderName: %s | isDir: %t \n", folderName, isDir)
				pathFoldersInsidePerUser := pathPerUser + "/" + folderName
				//folder path
				//user folder
				service.ExtractDataPerFile(pathFoldersInsidePerUser, folderName, userFolder)

			}
		}
	}

}
