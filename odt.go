// Copyright © 2023 Mark Summerfield. All rights reserved.
// License: Apache-2.0

package odt

import (
	"archive/zip"
	_ "embed"
	"io"
)

//go:embed Version.dat
var Version string

type Odt struct {
	Filename string
	Files    []File
}

type File struct {
	Name string
	Text string
}

func Open(filename string) (*Odt, error) {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	odt := &Odt{Filename: filename, Files: make([]File, 0)}
	for _, file := range reader.File {
		itemReader, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer itemReader.Close()
		raw, err := io.ReadAll(itemReader)
		if err != nil {
			return nil, err
		}
		odt.Files = append(odt.Files, File{Name: file.Name,
			Text: string(raw)})
	}
	return odt, nil
}
