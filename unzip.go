package mwbparser

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func unzip(sourcePath string, destinationPath string) error {
	zipFile, err := zip.OpenReader(sourcePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := zipFile.Close(); err != nil {
			panic(err)
		}
	}()

	err = os.MkdirAll(destinationPath, 0755)
	if err != nil {
		return err
	}

	extractAndWriteFile := func(file *zip.File) error {
		if strings.HasPrefix(file.Name, "@") {
			return nil
		}
		if filepath.Ext(file.Name) != ".xml" {
			return nil
		}
		readContext, err := file.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := readContext.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(destinationPath, file.Name)

		if file.FileInfo().IsDir() {
			err = os.MkdirAll(path, file.Mode())
			if err != nil {
				return err
			}
		} else {
			err = os.MkdirAll(filepath.Dir(path), file.Mode())
			if err != nil {
				return err
			}
			targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := targetFile.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(targetFile, readContext)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, file := range zipFile.File {
		if file.FileInfo().IsDir() {
			fmt.Println("Directory")
		}
		err := extractAndWriteFile(file)
		if err != nil {
			return err
		}
	}

	return nil
}
