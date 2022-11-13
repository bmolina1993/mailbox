package services

import (
	"regexp"
	"strings"

	entity "github.com/bmolina1993/mailbox/src/entities"
)

// extrae dato de propiedades del correo y contenido
func ExtractData(data string) (props, body string) {
	dataSplit := strings.Split(data, "X-FileName:")

	//obtiene todas las propiedades iniciales antes de [X-FileName]
	properties := dataSplit[0]
	//fmt.Println(properties)

	//obtiene todo el contenido luego de [X-FileName]
	auxBody := dataSplit[1]

	//obtiene todo el contenido, luego del 1er salto de linea
	auxBodyContent := strings.SplitN(auxBody, "\n", 2)
	bodyContent := auxBodyContent[1]

	// retorna propiedades y body
	return properties, bodyContent
}

// extrae fecha de envio
func extractPropDate(props string) string {
	pattern := regexp.MustCompile("(Date:)(.*)")
	matchSubstring := pattern.FindString(props)

	//obtiene contenido despues de [Date:]
	date := pattern.ReplaceAllString(matchSubstring, "$2")
	//quita espacios laterales
	date = strings.TrimSpace(date)

	return date
}

func extractPropFrom(props string) string {
	pattern := regexp.MustCompile("(From:)(.*)")
	matchSubstring := pattern.FindString(props)

	//obtiene contenido despues de [From:]
	from := pattern.ReplaceAllString(matchSubstring, "$2")
	//quita espacios laterales
	from = strings.TrimSpace(from)

	return from
}

// extrae los destinatarios del correo
// pagina prueba regex: https://regex101.com/
func extractPropTo(props string) []string {
	//obtiene contenido entre valores [To:] y [Subject:]
	pattern := regexp.MustCompile("(To:)((.|\n)*)(Subject:)")
	matchSubstring := pattern.FindString(props)

	//obtiente 2do grupo de datos
	//contenido entre los 2 valor declarados anteriormente
	groupExtract := pattern.ReplaceAllString(matchSubstring, "$2")

	//extrae cada correo
	mails := strings.Split(groupExtract, ",")

	//quita espacios laterales de cada elemento
	for i := range mails {
		mails[i] = strings.TrimSpace(mails[i])
	}

	return mails
}

func extractPropSubject(props string) string {
	pattern := regexp.MustCompile("(Subject:)((.|\n)*)(Cc:|Mime-Version:)")
	matchSubstring := pattern.FindString(props)

	//obtiene contenido despues de [Subject:]
	auxData := pattern.ReplaceAllString(matchSubstring, "$2")
	//si existe [Cc:], obtiene 1er indice para extraer data
	data := strings.Split(auxData, "Cc:")
	//quita espacios laterales
	subject := strings.TrimSpace(data[0])

	return subject
}

// extra data de cada archivo
// por carpeta interna de usuarios
func ExtractDataPerFile(path, folder, userFolder string) (a int, b map[string]string) {
	files, _ := ReadDirFile(path)

	//instanciamos struct mail
	dataMail := entity.QueryMail{}
	dataFolders := make(map[string]string)
	var qtyFolders int
	var qtyFiles int
	//var folderNameNested string

	//recorre cada archivo
	for fileName, isDir := range files {
		fullPathFile := path + "/" + fileName
		//extrae data solo de archivos
		if !isDir {
			//acumula cantidad de archivos por carpeta
			qtyFiles++
			//extraccion de dato mail
			auxData, _ := ReadFile(fullPathFile)
			props, body := ExtractData(auxData)

			date := extractPropDate(props)
			from := extractPropFrom(props)
			to := extractPropTo(props)
			subject := extractPropSubject(props)

			dataMail.Index = userFolder
			dataMail.Records = append(dataMail.Records, entity.Properties{
				Folder:  folder,
				Date:    date,
				From:    from,
				To:      to,
				Subject: subject,
				Body:    body,
			})
		} else {
			//si es carpeta, se acumula cantidad encontrada
			qtyFolders++
			dataFolders[fullPathFile] = folder + "/" + fileName
		}
	}

	//post api para insercion de datos masivo
	PostApiBulkData(dataMail)

	return qtyFolders, dataFolders
}
