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
