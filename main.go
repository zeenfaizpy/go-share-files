package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	_ "github.com/mdp/qrterminal"
)

func main() {
	//qrterminal.Generate("https://github.com/mdp/qrterminal", qrterminal.L, os.Stdout)

	argpath := os.Args[1]
	if argpath == "" {
		argpath, _ = os.Getwd()
	}

	zipfile, _ := os.Create("Result.zip")
	defer zipfile.Close()

	zipwriter := zip.NewWriter(zipfile)
	defer zipwriter.Close()

	info, err := zipfile.Stat()
	if err != nil {
		fmt.Println(err)
	}

	baseDir := filepath.Base(argpath)
	fileinfos, _ := ioutil.ReadDir(argpath)

	for _, fileinfo := range fileinfos {
		if !fileinfo.IsDir() {
			file_location := path.Join(argpath, fileinfo.Name())
			file, err := os.Open(file_location)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				fmt.Println(err)
			}
			destination := filepath.Join(baseDir, fileinfo.Name())
			header.Name = destination
			fmt.Printf("Archiving %s", destination)
			header.Method = zip.Deflate

			writer, _ := zipwriter.CreateHeader(header)
			_, err = io.Copy(writer, file)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
