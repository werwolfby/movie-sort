package main

import (
	"os"
	"path"
	"path/filepath"
)

type fileInfo struct {
	Folder     string      `json:"folder"`
	Path       []string    `json:"path"`
	Name       string      `json:"name"`
	OsFileInfo os.FileInfo `json:"-"`
}

type readDirFunc func(dirname string) ([]os.FileInfo, error)

type sameFileFunc func(fi1, fi2 os.FileInfo) bool

func (f *fileInfo) getFullName(folder folderInfo) string {
	return path.Join(append(append(folder.Path, f.Path...), f.Name)...)
}

func getAllFiles(reader readDirFunc, folderName string, folderPath []string, dirPath []string, extensions []string) ([]fileInfo, error) {
	dir := path.Join(append(folderPath, dirPath...)...)
	innerFiles, err := reader(dir)
	if err != nil {
		return nil, err
	}
	files := []fileInfo{}
	for _, f := range innerFiles {
		if !f.IsDir() {
			ext := filepath.Ext(f.Name())
			if len(ext) == 0 {
				continue
			}
			ext = ext[1:]
			for _, e := range extensions {
				if e == ext {
					files = append(files, fileInfo{folderName, dirPath, f.Name(), f})
					break
				}
			}
		} else {
			dirFiles, dirErr := getAllFiles(reader, folderName, folderPath, append(dirPath, f.Name()), extensions)
			if dirErr == nil {
				files = append(files, dirFiles...)
			}
		}
	}
	return files, nil
}
