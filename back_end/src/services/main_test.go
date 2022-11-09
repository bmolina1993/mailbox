package services

import (
	"testing"
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

		if err != nil {
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
