package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testDrive2 = mockDrive{
	"D:/Downloads/Complete":                 []string{"1.mkv", "Arrow.S01E08.rus.LostFilm.TV.avi", "Arrow.S01E09.rus.LostFilm.TV.avi", "Пианистка DVDRip.avi"},
	"D:/Downloads/Complete/TBBT (S09) 720p": []string{"TBBT.S09E01.HDTV.720p.KB.[qqss44].mkv", "TBBT.S09E02.HDTV.720p.KB.[qqss44].mkv"},
	"D:/Video/Shows/Arrow/Season 1":         []string{"Arrow.S01E08.rus.LostFilm.TV.avi", "Arrow.S01E09.rus.LostFilm.TV.avi"},
	"D:/Video/Shows/TBBT/Season 9":          []string{"TBBT.S09E01.HDTV.720p.KB.[qqss44].mkv", "TBBT.S09E02.HDTV.720p.KB.[qqss44].mkv"},
	"D:/Video/Movies":                       []string{"Пианистка DVDRip.avi"},
}

func sameFile(fi1, fi2 os.FileInfo) bool {
	return fi1.Name() == fi2.Name()
}

func TestSearchLinks(t *testing.T) {
	readDir := func(dirname string) ([]os.FileInfo, error) {
		return testDrive2.getFiles(dirname), nil
	}

	result := searchHardLinks(readDir, sameFile,
		[]folderInfo{folderInfo{"Downloads", []string{"D:", "Downloads", "Complete"}}},
		[]folderInfo{folderInfo{"Shows", []string{"D:", "Video", "Shows"}},
			folderInfo{"Movies", []string{"D:", "Video", "Movies"}}})

	assert.NotNil(t, result)

	expected := []linkInfo{
		{fileInfo: fileInfo{"Downloads", []string{"TBBT (S09) 720p"}, "TBBT.S09E01.HDTV.720p.KB.[qqss44].mkv", nil}, links: []fileInfo{
			fileInfo{"Shows", []string{"TBBT", "Season 9"}, "TBBT.S09E01.HDTV.720p.KB.[qqss44].mkv", nil},
		}},
		{fileInfo: fileInfo{"Downloads", []string{"TBBT (S09) 720p"}, "TBBT.S09E02.HDTV.720p.KB.[qqss44].mkv", nil}, links: []fileInfo{
			fileInfo{"Shows", []string{"TBBT", "Season 9"}, "TBBT.S09E02.HDTV.720p.KB.[qqss44].mkv", nil},
		}},
		{fileInfo: fileInfo{"Downloads", []string{}, "1.mkv", nil}, links: nil},
		{fileInfo: fileInfo{"Downloads", []string{}, "Arrow.S01E08.rus.LostFilm.TV.avi", nil}, links: []fileInfo{
			fileInfo{"Shows", []string{"Arrow", "Season 1"}, "Arrow.S01E08.rus.LostFilm.TV.avi", nil},
		}},
		{fileInfo: fileInfo{"Downloads", []string{}, "Arrow.S01E09.rus.LostFilm.TV.avi", nil}, links: []fileInfo{
			fileInfo{"Shows", []string{"Arrow", "Season 1"}, "Arrow.S01E09.rus.LostFilm.TV.avi", nil},
		}},
		{fileInfo: fileInfo{"Downloads", []string{}, "Пианистка DVDRip.avi", nil}, links: []fileInfo{
			fileInfo{"Movies", []string{}, "Пианистка DVDRip.avi", nil},
		}},
	}

	for i, fi := range expected {
		fillFileInfo(&fi.fileInfo)
		for j, li := range fi.links {
			fillFileInfo(&li)
			fi.links[j] = li
		}
		expected[i] = fi
	}

	assert.Equal(t, expected, result)
}
