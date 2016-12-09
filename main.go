package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/archiver/extractor"
	"github.com/dpastoor/nonmemutils/utils"
	"github.com/spf13/afero"
)

func main() {
	zipE := extractor.NewZip()
	AppFs := afero.NewOsFs()
	dir := filepath.Dir(".")
	dirInfo, _ := afero.ReadDir(AppFs, dir)
	files := utils.ListFiles(dirInfo)
	fmt.Println(files)
	for _, file := range files {
		fullPath := filepath.Join(dir, file)
		fileName, _ := utils.FileAndExt(fullPath)
		zipE.Extract(fullPath, filepath.Join(dir, fileName))
	}
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
