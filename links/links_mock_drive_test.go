package links_test

import (
	"os"
	"strings"
	"time"
)

type MockFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (fi *MockFileInfo) Name() string {
	return fi.name
}

func (fi *MockFileInfo) Size() int64 {
	return fi.size
}

func (fi *MockFileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi *MockFileInfo) ModTime() time.Time {
	return fi.modTime
}

func (fi *MockFileInfo) IsDir() bool {
	return fi.isDir
}

func (fi *MockFileInfo) Sys() interface{} {
	return nil
}

type mockDrive map[string][]string

func (drive mockDrive) getFiles(name string) []os.FileInfo {
	f := drive[name]
	if f == nil {
		f = []string{}
	}

	subfolders := []string{}
	if drive != nil {
		for k := range drive {
			if strings.HasPrefix(k, name+"\\") {
				subfolder := k[len(name)+1:]
				index := strings.Index(subfolder, "\\")
				if index > 0 {
					subfolder = subfolder[:strings.Index(subfolder, "\\")]
				}
				subfolders = append(subfolders, subfolder)
			}
		}
	}

	result := []os.FileInfo{}

	for _, f := range subfolders {
		result = append(result, &MockFileInfo{name: f, isDir: true})
	}

	for _, f := range f {
		result = append(result, &MockFileInfo{name: f})
	}

	return result
}

func (drive mockDrive) ReadDir(dirname string) ([]os.FileInfo, error) {
	return drive.getFiles(dirname), nil
}

func (drive mockDrive) SameFile(fi1, fi2 os.FileInfo) bool {
	return fi1.Name() == fi2.Name()
}

func (mockDrive) MkdirAll(dirname string) error {
	return nil
}

func (drive mockDrive) Link(oldname, newname string) error {
	return nil
}
