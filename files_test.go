package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
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

type mockFolder struct {
	folder string
	files  []string
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
			if strings.HasPrefix(k, name+"/") {
				subfolder := k[len(name)+1:]
				index := strings.Index(subfolder, "/")
				if index > 0 {
					subfolder = subfolder[:strings.Index(subfolder, "/")]
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

var testDrive1 = mockDrive{
	"D:/Downloads/Complete":                 []string{"1.mkv", "Arrow.S01E08.rus.LostFilm.TV.avi", "Arrow.S01E09.rus.LostFilm.TV.avi", "Пианистка DVDRip.avi"},
	"D:/Downloads/Complete/TBBT (S09) 720p": []string{"TBBT.S09E01.HDTV.720p.KB.[qqss44].mkv", "TBBT.S09E02.HDTV.720p.KB.[qqss44].mkv"},
}

func TestGetAllFiles(t *testing.T) {
	readDir := func(dirname string) ([]os.FileInfo, error) {
		return testDrive1.getFiles(dirname), nil
	}

	files, err := getAllFiles(readDir, "Downloads", []string{"D:", "Downloads", "Complete"}, []string{}, []string{"mkv", "avi"})

	expected := []fileInfo{
		fileInfo{"Downloads", []string{"TBBT (S09) 720p"}, "TBBT.S09E01.HDTV.720p.KB.[qqss44].mkv", nil},
		fileInfo{"Downloads", []string{"TBBT (S09) 720p"}, "TBBT.S09E02.HDTV.720p.KB.[qqss44].mkv", nil},
		fileInfo{"Downloads", []string{}, "1.mkv", nil},
		fileInfo{"Downloads", []string{}, "Arrow.S01E08.rus.LostFilm.TV.avi", nil},
		fileInfo{"Downloads", []string{}, "Arrow.S01E09.rus.LostFilm.TV.avi", nil},
		fileInfo{"Downloads", []string{}, "Пианистка DVDRip.avi", nil},
	}

	for i, fi := range expected {
		fillFileInfo(&fi)
		expected[i] = fi
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, files)
}

func fillFileInfo(f *fileInfo) {
	f.OsFileInfo = &MockFileInfo{name: f.Name}
}
