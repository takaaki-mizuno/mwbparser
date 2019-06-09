package mwbparser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUnzip(t *testing.T) {
	absolutePath, err := filepath.Abs("./data/db.mwb")
	tempDirectory := os.TempDir()
	destinationPath := filepath.Join(tempDirectory, "test_unzip_mwbparser")
	err = unzip(absolutePath, destinationPath)
	if err != nil {
		t.Fatalf("failed to unzip %#v", err)
	}
	path := filepath.Join(destinationPath, "document.mwb.xml")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("cannot find document.mwb.xml %#v", err)
	}

	defer func() {
		if err := os.RemoveAll(destinationPath); err != nil {
			panic(err)
		}
	}()
}
