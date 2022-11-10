package services

import (
	"reflect"
	"testing"

	entity "github.com/bmolina1993/mailbox/src/entities"
)

// valida si en la ruta ingresada
// se encuentran carpetas o archivos
func TestReadDirFile(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 3

	tables := []struct {
		path   string
		folder string
		isDir  bool
	}{
		{
			"../../data",
			"allen-p",
			true,
		},
		{
			"../../data/badeer-r/inbox",
			"1.",
			false,
		},
		{
			"../../data/badeer-r/inboxError",
			"Error",
			false,
		},
	}

	for _, item := range tables {
		operation, err := ReadDirFile(item.path)

		if err != "" {
			qtyPass++
		}

		for name, isDir := range operation {
			if name == item.folder && isDir == item.isDir {
				qtyPass++
			}
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestReadFile(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2

	tables := []struct {
		pathFile string
		dataMail string
	}{
		{
			"../../data/allen-p/inbox/1.",
			"",
		},
		{
			"../../data/allen-p/inbox/ERROR",
			"",
		},
	}

	for _, item := range tables {
		operation, err := ReadFile(item.pathFile)

		if err != nil {
			qtyPass++
		}

		if operation == item.dataMail {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

// extrae dato de propiedades del correo y contenido
func TestExtractData(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		pathFile string
		props    string
		body     string
	}{
		{
			"../../data/allen-p/inbox/1.",
			"",
			"",
		},
		{
			"../../data/allen-p/inbox/2.",
			"",
			"",
		},
	}

	for _, item := range tables {
		data, _ := ReadFile(item.pathFile)
		props, body := ExtractData(data)

		if props != item.props && body != item.body {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

// extrae fecha de envio
func TestExtractPropDate(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		pathFile string
		date     string
	}{
		{
			"../../data/allen-p/inbox/1.",
			"",
		},
		{
			"../../data/allen-p/inbox/2.",
			"",
		},
	}

	for _, item := range tables {
		data, _ := ReadFile(item.pathFile)
		props, _ := ExtractData(data)
		date := extractPropDate(props)

		if date != item.date {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestExtractPropFrom(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		pathFile string
		from     string
	}{
		{
			"../../data/allen-p/inbox/1.",
			"",
		},
		{
			"../../data/allen-p/inbox/2.",
			"",
		},
	}

	for _, item := range tables {
		data, _ := ReadFile(item.pathFile)
		props, _ := ExtractData(data)
		from := extractPropFrom(props)

		if from != item.from {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestExtractPropTo(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		pathFile string
		to       []string
	}{
		{
			"../../data/allen-p/inbox/1.",
			[]string{},
		},
		{
			"../../data/allen-p/inbox/2.",
			[]string{},
		},
	}

	for _, item := range tables {
		data, _ := ReadFile(item.pathFile)
		props, _ := ExtractData(data)
		to := extractPropTo(props)

		if len(to) != len(item.to) {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestExtractPropSubject(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		pathFile string
		subject  string
	}{
		{
			"../../data/allen-p/inbox/1.",
			"",
		},
		{
			"../../data/allen-p/inbox/2.",
			"",
		},
	}

	for _, item := range tables {
		data, _ := ReadFile(item.pathFile)
		props, _ := ExtractData(data)
		subject := extractPropSubject(props)

		if subject != item.subject {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestExtractDataPerFile(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		path       string
		folder     string
		userFolder string

		qtyFolders  int
		dataFolders map[string]string
	}{
		{
			"../../data/allen-p/",
			"inbox",
			"allen-p",
			10,
			map[string]string{},
		},
		{
			"../../data/baughman-d/",
			"inbox",
			"baughman-d",
			15,
			map[string]string{},
		},
	}

	for _, item := range tables {
		qtyFolders, dataFolders := ExtractDataPerFile(item.path, item.folder, item.userFolder)

		if qtyFolders == item.qtyFolders && !reflect.DeepEqual(dataFolders, item.dataFolders) {
			qtyPass++
		}
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestExtractAllFolders(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 2
	tables := []struct {
		pathFile string
		dirFile  map[string]bool
	}{
		{
			"../../data/",
			map[string]bool{
				"stokley-c": true,
			},
		},
		{
			"../../data/",
			map[string]bool{
				"motley-m": true,
			},
		},
		{
			"../../dataERROR",
			map[string]bool{
				"motley-m": true,
			},
		},
	}

	for _, item := range tables {
		err := ExtractAllFolders(item.pathFile, item.dirFile)

		if err == "" {
			qtyPass++
		}

	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestPostApiBulkData(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 1

	err := PostApiBulkData(entity.QueryMail{Index: "testing", Records: []entity.Properties{}})
	if err == nil {
		qtyPass++
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestGetIdsDocuments(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 1

	_, err := GetIdsDocuments()
	if err == nil {
		qtyPass++
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}

func TestDeleteDocuments(t *testing.T) {
	var qtyPass int
	qtyPassExpected := 1

	err := DeleteDocuments("stokley-c")
	if err == nil {
		qtyPass++
	}

	if qtyPass != qtyPassExpected {
		t.Errorf("Se esperaba %d pase/s y se obtuvieron %d", qtyPassExpected, qtyPass)
	}
}
