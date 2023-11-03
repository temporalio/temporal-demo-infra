package app

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type justFilesFilesystem struct {
	fs http.FileSystem
	// readDirBatchSize - configuration parameter for `Readdir` func
	readDirBatchSize int
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	if name[len(name)-1] != '/' && filepath.Ext(name) == "" {
		name = name + ".html"
	}
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredStatFile{File: f, readDirBatchSize: fs.readDirBatchSize}, nil
}

type neuteredStatFile struct {
	http.File
	readDirBatchSize int
}

func (e neuteredStatFile) Stat() (os.FileInfo, error) {
	s, err := e.File.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
	LOOP:
		for {
			fl, ferr := e.File.Readdir(e.readDirBatchSize)
			switch ferr {
			case io.EOF:
				break LOOP
			case nil:
				for _, f := range fl {
					if f.Name() == "index.html" {
						return s, err
					}
				}
			default:
				return nil, ferr
			}
		}
		return nil, os.ErrNotExist
	}
	return s, err
}
