package main

import (
	service "back_end/src/services"
	"fmt"
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
		if fileToDelete != "shively-h" {
			delete(dirFile, fileToDelete)
		}
	}
	fmt.Println("--> filtra Carpeta <--")
	fmt.Println(dirFile)
	// TEMPORAL [X]

	/**/
	fmt.Println("--> Recorre carpeta por usuario <--")

	//accede a carpeta de usuarios
	for userFolder := range dirFile {
		// iteracion por usuario,
		// se omiten archivos por fuera de carpetas
		dirPerUser := service.ReadDirFile(pathFile + userFolder)
		for fileName, isDir := range dirPerUser {
			if isDir {
				fmt.Printf("fileName: %s | isDir: %t \n", fileName, isDir)
			}
		}
	}

}
