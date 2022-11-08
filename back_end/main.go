package main

import (
	"fmt"

	service "github.com/bmolina1993/mailbox/src/services"
)

func main() {
	pathFile := "./data/"

	//caso de usuario con archivo suelto en carpeta usuario
	//dirFile := ReadDirFile("./data/shively-h/")
	dirFile := service.ReadDirFile(pathFile)
	fmt.Println("--> Carpetas <--")
	fmt.Println(dirFile)

	// TEMPORAL [X]
	// filtra carpeta
	for fileToDelete := range dirFile {
		//motley-m -> 4 niveles | stokley-c -> 6 niveles
		if fileToDelete != "stokley-c" {
			delete(dirFile, fileToDelete)
		}
	}
	fmt.Println("\n--> filtra Carpeta <--")
	fmt.Println(dirFile)
	// TEMPORAL [X]

	fmt.Println("\n--> Recorre carpeta por usuario <--")

	//accede a carpeta de usuarios
	for userFolder := range dirFile {
		// iteracion por usuario,
		// se omiten archivos por fuera de carpetas
		pathPerUser := pathFile + userFolder
		dirPerUser := service.ReadDirFile(pathPerUser)

		// TEMPORAL [X]
		// filtra carpeta
		for folderToDelete := range dirPerUser { // [1/6 niveles]
			//motley-m:inbox -> carpetas anidadas | motley-m:sent_items -> sin carpetas dentro
			//stokley-c:chris_stokley -> carpetas anidadas
			if folderToDelete != "chris_stokley" {
				delete(dirPerUser, folderToDelete)
			}
		}
		fmt.Println("\n--> filtra Carpeta interna de usuario <--")
		fmt.Println(dirPerUser)
		// TEMPORAL [X]

		fmt.Println("\n--> recorre cada archivo por carpeta de usuario <--")
		for folderName, isDir := range dirPerUser { // [2/6 niveles]
			if isDir {
				pathFoldersInsidePerUser := pathPerUser + "/" + folderName
				qtyFolders, dataFolders := service.ExtractDataPerFile(pathFoldersInsidePerUser, folderName, userFolder)

				for qtyFolders > 0 { // [3/6 niveles]
					for fullPathFile, folderNested := range dataFolders {

						fmt.Println("\n--> ExtractDataPerFile [PARAMS] [3/6] <--")
						fmt.Println("fullPathFile:", fullPathFile)
						fmt.Println("folderNested:", folderNested)
						fmt.Println("userFolder:", userFolder)

						qtyFolders, dataFolders = service.ExtractDataPerFile(fullPathFile, folderNested, userFolder)

						for qtyFolders > 0 { // [4/6 niveles]
							for fullPathFile, folderNested := range dataFolders {

								fmt.Println("\n--> ExtractDataPerFile [PARAMS] [4/6] <--")
								fmt.Println("fullPathFile:", fullPathFile)
								fmt.Println("folderNested:", folderNested)
								fmt.Println("userFolder:", userFolder)

								qtyFolders, dataFolders = service.ExtractDataPerFile(fullPathFile, folderNested, userFolder)

								for qtyFolders > 0 { // [5/6 niveles]
									for fullPathFile, folderNested := range dataFolders {

										fmt.Println("\n--> ExtractDataPerFile [PARAMS] [5/6] <--")
										fmt.Println("fullPathFile:", fullPathFile)
										fmt.Println("folderNested:", folderNested)
										fmt.Println("userFolder:", userFolder)

										qtyFolders, dataFolders = service.ExtractDataPerFile(fullPathFile, folderNested, userFolder)

										for qtyFolders > 0 { // [6/6 niveles]
											for fullPathFile, folderNested := range dataFolders {

												fmt.Println("\n--> ExtractDataPerFile [PARAMS] [6/6] <--")
												fmt.Println("fullPathFile:", fullPathFile)
												fmt.Println("folderNested:", folderNested)
												fmt.Println("userFolder:", userFolder)

												qtyFolders, dataFolders = service.ExtractDataPerFile(fullPathFile, folderNested, userFolder)

												fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [6/6] <--")
												fmt.Println("qtyFolders:", qtyFolders)
												fmt.Println("dataFolders:", dataFolders)
											}
										}

										fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [5/6] <--")
										fmt.Println("qtyFolders:", qtyFolders)
										fmt.Println("dataFolders:", dataFolders)
									}
								}

								fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [4/6] <--")
								fmt.Println("qtyFolders:", qtyFolders)
								fmt.Println("dataFolders:", dataFolders)
							}
						}

						fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [3/6] <--")
						fmt.Println("qtyFolders:", qtyFolders)
						fmt.Println("dataFolders:", dataFolders)

					}
				}
			}
		}
	}

}
