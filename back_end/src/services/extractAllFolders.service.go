package services

func ExtractAllFolders(pathFile string, dirFile map[string]bool) string {
	var errMsg string
	/*
		// TEMPORAL [X]
		// filtra carpeta
		for fileToDelete := range dirFile {
			//motley-m -> 4 niveles | "stokley-c" -> 6 niveles
			if fileToDelete != "stokley-c" {
				delete(dirFile, fileToDelete)
			}
		}
		fmt.Println("\n--> filtra Carpeta <--")
		fmt.Println(dirFile)
		// TEMPORAL [X]
	*/

	//accede a carpeta de usuarios
	for userFolder := range dirFile { // [1/6 niveles]
		// iteracion por usuario,
		// se omiten archivos por fuera de carpetas
		pathPerUser := pathFile + userFolder
		dirPerUser, err := ReadDirFile(pathPerUser)
		if err != "" {
			errMsg = err
		}

		/*
			// TEMPORAL [X]
			// filtra carpeta
			for folderToDelete := range dirPerUser {
				//motley-m:inbox -> carpetas anidadas | motley-m:sent_items -> sin carpetas dentro
				//stokley-c:chris_stokley -> carpetas anidadas
				if folderToDelete != "chris_stokley" {
					delete(dirPerUser, folderToDelete)
				}
			}
			fmt.Println("\n--> filtra sub-carpeta de usuario <--")
			fmt.Println(dirPerUser)
			// TEMPORAL [X]
		*/

		for folderName, isDir := range dirPerUser { // [2/6 niveles]
			//fmt.Println("\n--> recorre cada archivo por carpeta de usuario [2/6] <--")
			if isDir {
				pathFoldersInsidePerUser := pathPerUser + "/" + folderName
				qtyFolders, dataFolders := ExtractDataPerFile(pathFoldersInsidePerUser, folderName, userFolder)

				for qtyFolders > 0 { // [3/6 niveles]
					for fullPathFile, folderNested := range dataFolders {

						/*
							fmt.Println("\n--> ExtractDataPerFile [PARAMS] [3/6] <--")
							fmt.Println("fullPathFile:", fullPathFile)
							fmt.Println("folderNested:", folderNested)
							fmt.Println("userFolder:", userFolder)
						*/

						qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

						for qtyFolders > 0 { // [4/6 niveles]
							for fullPathFile, folderNested := range dataFolders {

								/*
									fmt.Println("\n--> ExtractDataPerFile [PARAMS] [4/6] <--")
									fmt.Println("fullPathFile:", fullPathFile)
									fmt.Println("folderNested:", folderNested)
									fmt.Println("userFolder:", userFolder)
								*/

								qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

								for qtyFolders > 0 { // [5/6 niveles]
									for fullPathFile, folderNested := range dataFolders {

										/*
											fmt.Println("\n--> ExtractDataPerFile [PARAMS] [5/6] <--")
											fmt.Println("fullPathFile:", fullPathFile)
											fmt.Println("folderNested:", folderNested)
											fmt.Println("userFolder:", userFolder)
										*/

										qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

										for qtyFolders > 0 { // [6/6 niveles]
											for fullPathFile, folderNested := range dataFolders {

												/*
													fmt.Println("\n--> ExtractDataPerFile [PARAMS] [6/6] <--")
													fmt.Println("fullPathFile:", fullPathFile)
													fmt.Println("folderNested:", folderNested)
													fmt.Println("userFolder:", userFolder)
												*/

												qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

												/*
													fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [6/6] <--")
													fmt.Println("qtyFolders:", qtyFolders)
													fmt.Println("dataFolders:", dataFolders)
												*/
											}
										}

										/*
											fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [5/6] <--")
											fmt.Println("qtyFolders:", qtyFolders)
											fmt.Println("dataFolders:", dataFolders)
										*/
									}
								}

								/*
									fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [4/6] <--")
									fmt.Println("qtyFolders:", qtyFolders)
									fmt.Println("dataFolders:", dataFolders)
								*/
							}
						}

						/*
							fmt.Println("\n--> ExtractDataPerFile [RETURNS DATA] [3/6] <--")
							fmt.Println("qtyFolders:", qtyFolders)
							fmt.Println("dataFolders:", dataFolders)
						*/

					}
				}
			}
		}
	}

	return errMsg
}
