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
	Files    map[string]string // key=filename value=text (usually XML)
}

func Open(filename string) (*Odt, error) {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	odt := &Odt{Filename: filename, Files: make(map[string]string)}
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
		odt.Files[file.Name] = string(raw)
	}
	return odt, nil
}
