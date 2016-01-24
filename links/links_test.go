package links_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/werwolfby/movie-sort/links"
	"github.com/werwolfby/movie-sort/settings"
	"testing"
)

func TestUpdateLinks(t *testing.T) {
	drive := mockDrive{
		"D:\\Torrents":                  []string{"1.mkv", "2.avi", "3.mp4", "4.jpg"},
		"D:\\Movies":                    []string{"1.mkv"},
		"D:\\Shows\\The Show\\Season 3": []string{"2.avi"},
	}

	inputFolders := []settings.FolderInfo{settings.FolderInfo{Name: "Downloads", Path: []string{"D:", "Torrents"}}}
	outputFolders := []settings.FolderInfo{
		settings.FolderInfo{Name: "Movies", Path: []string{"D:", "Movies"}, Meta: settings.FolderMetaMovies},
		settings.FolderInfo{Name: "Shows", Path: []string{"D:", "Shows"}, Meta: settings.FolderMetaShows},
	}

	ifs := settings.InputFoldersSettings{FoldersSettings: settings.FoldersSettings{Folders: inputFolders}}
	ofs := settings.OutputFoldersSettings{FoldersSettings: settings.FoldersSettings{Folders: outputFolders}}
	l := links.NewLinks(drive, &ifs, &ofs)

	l.UpdateLinks([]string{"mkv", "avi", "mp4"})

	ls := l.GetLinks()

	assert.NotNil(t, ls)

	expectedLinks := []links.LinkInfo{
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "1.mkv"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Movies", Path: nil, Name: "1.mkv"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "2.avi"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Shows", Path: []string{"The Show", "Season 3"}, Name: "2.avi"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "3.mp4"}, Links: nil},
	}

	assert.Equal(t, expectedLinks, ls)

	expectedShows := []links.FileInfo{
		links.FileInfo{Folder: "Shows", Path: []string{}, Name: "The Show"},
	}

	assert.Equal(t, expectedShows, l.GetShows())
}

func TestUpdateLinks2(t *testing.T) {
	drive := mockDrive{
		"D:\\Torrents":                  []string{"1.mkv", "The.Show.S03E01.avi", "The.Show.S03E02.avi", "3.mp4", "4.jpg"},
		"D:\\Movies":                    []string{"1.mkv"},
		"D:\\Shows\\The Show\\Season 3": []string{"The.Show.S03E01.avi"},
		"D:\\Shows\\the show\\Season 3": []string{"The.Show.S03E02.avi"}, // For list of shows case have to be ignored
		"D:\\Shows\\Other Show":         []string{"3.mp4"},
	}

	inputFolders := []settings.FolderInfo{settings.FolderInfo{Name: "Downloads", Path: []string{"D:", "Torrents"}}}
	outputFolders := []settings.FolderInfo{
		settings.FolderInfo{Name: "Movies", Path: []string{"D:", "Movies"}, Meta: settings.FolderMetaMovies},
		settings.FolderInfo{Name: "Shows", Path: []string{"D:", "Shows"}, Meta: settings.FolderMetaShows},
	}

	ifs := settings.InputFoldersSettings{FoldersSettings: settings.FoldersSettings{Folders: inputFolders}}
	ofs := settings.OutputFoldersSettings{FoldersSettings: settings.FoldersSettings{Folders: outputFolders}}
	l := links.NewLinks(drive, &ifs, &ofs)

	l.UpdateLinks([]string{"mkv", "avi", "mp4"})

	ls := l.GetLinks()

	assert.NotNil(t, ls)

	expectedLinks := []links.LinkInfo{
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "1.mkv"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Movies", Path: nil, Name: "1.mkv"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "The.Show.S03E01.avi"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Shows", Path: []string{"The Show", "Season 3"}, Name: "The.Show.S03E01.avi"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "The.Show.S03E02.avi"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Shows", Path: []string{"the show", "Season 3"}, Name: "The.Show.S03E02.avi"},
		}},
		links.LinkInfo{FileInfo: links.FileInfo{Folder: "Downloads", Path: nil, Name: "3.mp4"}, Links: []links.FileInfo{
			links.FileInfo{Folder: "Shows", Path: []string{"Other Show"}, Name: "3.mp4"},
		}},
	}

	assert.Equal(t, expectedLinks, ls)

	expectedShows := []links.FileInfo{
		links.FileInfo{Folder: "Shows", Path: []string{}, Name: "The Show"},
		links.FileInfo{Folder: "Shows", Path: []string{}, Name: "Other Show"},
	}

	assert.Equal(t, expectedShows, l.GetShows())

	show, found := l.GetShow("ThE SHoW")
	assert.True(t, found)
	assert.Equal(t, expectedShows[0], show)

	show, found = l.GetShow("oThEr SHoW")
	assert.True(t, found)
	assert.Equal(t, expectedShows[1], show)

	show, found = l.GetShow("Unknown Show")
	assert.False(t, found)

	expectedSeasons := []links.FileInfo{
		links.FileInfo{Folder: "Shows", Path: []string{"Season 3"}, Name: "The Show"},
		expectedShows[1],
	}

	var season int
	show, season = l.GetShowSeason("ThE SHoW", 2)
	assert.Equal(t, -1, season)
	// Show folder without seasons
	assert.Equal(t, expectedShows[0], show)

	show, season = l.GetShowSeason("oThEr SHoW", 2)
	assert.Equal(t, -1, season)
	// Show folder without seasons
	assert.Equal(t, expectedShows[1], show)

	show, season = l.GetShowSeason("Unknown Show", 1)
	assert.Equal(t, -2, season)
	// Show folder without seasons
	assert.Equal(t, links.FileInfo{}, show)

	show, season = l.GetShowSeason("ThE SHoW", 3)
	assert.Equal(t, 3, season)
	assert.Equal(t, expectedSeasons[0], show)
}
