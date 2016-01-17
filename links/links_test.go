package links_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/werwolfby/movie-sort/links"
	"github.com/werwolfby/movie-sort/settings"
	"testing"
)

func TestUpdateLinks(t *testing.T) {
	drive1 := mockDrive{
		"D:\\Torrents":              []string{"1.mkv", "2.avi", "3.mp4", "4.jpg"},
		"D:\\Movies":                []string{"1.mkv"},
		"D:\\Shows\\Show\\Season 3": []string{"2.avi"},
	}

	inputFolders := []settings.FolderInfo{settings.FolderInfo{Name: "Downloads", Path: []string{"D:", "Torrents"}}}
	outputFolders := []settings.FolderInfo{
		settings.FolderInfo{Name: "Movies", Path: []string{"D:", "Movies"}},
		settings.FolderInfo{Name: "Shows", Path: []string{"D:", "Shows"}},
	}

	ifs := settings.InputFoldersSettings{FoldersSettings: settings.FoldersSettings{Folders: inputFolders}}
	ofs := settings.OutputFoldersSettings{FoldersSettings: settings.FoldersSettings{Folders: outputFolders}}
	l := links.NewLinks(drive1, &ifs, &ofs)

	l.UpdateLinks([]string{"mkv", "avi", "mp4"})

	ls := l.GetLinks()

	assert.NotNil(t, ls)

	expected := []links.LinkInfo{
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "1.mkv"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Movies", Path: nil, Name: "1.mkv"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "2.avi"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Shows", Path: []string{"Show", "Season 3"}, Name: "2.avi"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "3.mp4"}, Links: nil},
	}

	assert.Equal(t, expected, ls)
}
