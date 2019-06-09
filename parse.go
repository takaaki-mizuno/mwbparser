package mwbparser

import (
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/takaaki-mizuno/mwbparser/objects"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomDirectoryName(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) + "_" + fmt.Sprintf("%d", time.Now().UnixNano())
}

func Parse(mbwFilePath string) ([]objects.Table, error) {
	absolutePath, err := filepath.Abs(mbwFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Wrong path format: %s", mbwFilePath))
	}
	_, fileExistError := os.Stat(absolutePath)
	if fileExistError != nil {
		return nil, errors.New(fmt.Sprintf("MWB file didn't exist: %s", absolutePath))
	}
	tempDirectory := os.TempDir()
	destinationPath := filepath.Join(tempDirectory, randomDirectoryName(4))
	err = unzip(absolutePath, destinationPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(destinationPath); err != nil {
			panic(err)
		}
	}()

	xmlPath := filepath.Join(destinationPath, "document.mwb.xml")
	reader, err := os.Open(xmlPath)
	if err != nil {
		return nil, err
	}
	document, err := xmlquery.Parse(reader)

	tables, err := objects.ParseTables(document)
	if err != nil {
		return nil, err
	}

	return tables, nil
}
