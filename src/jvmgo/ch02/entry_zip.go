package main

import (
	"archive/zip"
	_ "archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)
import _ "errors"
import _ "io/ioutil"
import _ "path/filepath"

type zipEntry struct {
	absPath string
}

func newZipEntry(path string) *zipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &zipEntry{absPath}
}
func (self *zipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self *zipEntry) String() string {
	return self.absPath
}
