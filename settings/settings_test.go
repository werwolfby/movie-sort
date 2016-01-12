package settings_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/werwolfby/movie-sort/settings"
	"os"
	"strings"
	"testing"
)

var config1 = `
[paths]
src         = "/mnt/media/Torrents/Complete"
dest-movies = "/mnt/media/Video/Movies"
dest-shows  = "/mnt/media/Video/Shows"    
`

func TestReadSettingsPaths(t *testing.T) {
	r := strings.NewReader(config1)
	s, e := settings.ReadSettings(r)

	assert.Nil(t, e)
	assert.NotNil(t, s)

	assert.Equal(t, fmt.Sprintf("%c", os.PathSeparator), s.Global.PathSeparator)

	assert.Equal(t, 1, len(s.InputFolders.Folders))
	assert.Equal(t, 2, len(s.OutputFolders.Folders))

	assert.Equal(t, "Downloads", s.InputFolders.Folders[0].Name)
	assert.Equal(t, []string{"/mnt", "media", "Torrents", "Complete"}, s.InputFolders.Folders[0].Path)

	assert.Equal(t, "Movies", s.OutputFolders.Folders[0].Name)
	assert.Equal(t, []string{"/mnt", "media", "Video", "Movies"}, s.OutputFolders.Folders[0].Path)

	assert.Equal(t, "Shows", s.OutputFolders.Folders[1].Name)
	assert.Equal(t, []string{"/mnt", "media", "Video", "Shows"}, s.OutputFolders.Folders[1].Path)

	assert.Equal(t, "Downloads", s.FolderNames.DownloadsName)
	assert.Equal(t, "Movies", s.FolderNames.MoviesName)
	assert.Equal(t, "Shows", s.FolderNames.ShowsName)

	assert.Empty(t, s.Services.GuessItURL)
}

var config2 = `
[names]
downloads = "downloads"
movies    = "movies"
shows     = "shows"
`

func TestReadSettingsNames(t *testing.T) {
	r := strings.NewReader(config2)
	s, e := settings.ReadSettings(r)

	assert.Nil(t, e)
	assert.NotNil(t, s)

	assert.Equal(t, fmt.Sprintf("%c", os.PathSeparator), s.Global.PathSeparator)

	assert.Equal(t, 0, len(s.InputFolders.Folders))
	assert.Equal(t, 0, len(s.OutputFolders.Folders))

	assert.Equal(t, "downloads", s.FolderNames.DownloadsName)
	assert.Equal(t, "movies", s.FolderNames.MoviesName)
	assert.Equal(t, "shows", s.FolderNames.ShowsName)

	assert.Empty(t, s.Services.GuessItURL)
}

var config3 = `
[services]
guessit = "http://localhost:5000/guessit"
`

func TestReadSettingsServices(t *testing.T) {
	r := strings.NewReader(config3)
	s, e := settings.ReadSettings(r)

	assert.Nil(t, e)
	assert.NotNil(t, s)

	assert.Equal(t, fmt.Sprintf("%c", os.PathSeparator), s.Global.PathSeparator)

	assert.Equal(t, 0, len(s.InputFolders.Folders))
	assert.Equal(t, 0, len(s.OutputFolders.Folders))

	// Default values
	assert.Equal(t, "Downloads", s.FolderNames.DownloadsName)
	assert.Equal(t, "Movies", s.FolderNames.MoviesName)
	assert.Equal(t, "Shows", s.FolderNames.ShowsName)

	assert.Equal(t, "http://localhost:5000/guessit", s.Services.GuessItURL)
}
