package main

import (
	"fmt"
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
	for _, file := range files {
		fmt.Printf("unzipping %s \n", file)
		fullPath := filepath.Join(dir, file)
		fileName, _ := utils.FileAndExt(fullPath)
		zipE.Extract(fullPath, filepath.Join(dir, fileName))
	}
}
