package services

// lo niveles hace referencia a las subs-carpetas
// que puede acceder en terminos de profundidad
// este directorio, contiene hasta 6 niveles
// se obtiene por el siguiente comando: tree -R -d data/
func ExtractAllFolders(pathFile string, dirFile map[string]bool) error {

	//accede a carpeta de usuarios
	for userFolder := range dirFile { // [1/6 niveles]
		// iteracion por usuario,
		// se omiten archivos por fuera de carpetas
		pathPerUser := pathFile + userFolder
		dirPerUser, err := ReadDirFile(pathPerUser)
		if err != nil {
			return err
		}

		//recorre cada archivo por carpeta de usuario
		for folderName, isDir := range dirPerUser { // [2/6 niveles]
			if isDir {
				pathFoldersInsidePerUser := pathPerUser + "/" + folderName
				qtyFolders, dataFolders := ExtractDataPerFile(pathFoldersInsidePerUser, folderName, userFolder)

				for qtyFolders > 0 { // [3/6 niveles]
					for fullPathFile, folderNested := range dataFolders {
						qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

						for qtyFolders > 0 { // [4/6 niveles]
							for fullPathFile, folderNested := range dataFolders {
								qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

								for qtyFolders > 0 { // [5/6 niveles]
									for fullPathFile, folderNested := range dataFolders {
										qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)

										for qtyFolders > 0 { // [6/6 niveles]
											for fullPathFile, folderNested := range dataFolders {
												qtyFolders, dataFolders = ExtractDataPerFile(fullPathFile, folderNested, userFolder)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return nil
}
