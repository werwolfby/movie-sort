package settings

import (
	"fmt"
	"github.com/werwolfby/movie-sort/utils"
	"gopkg.in/gcfg.v1"
	"io"
	"os"
)

type configPaths struct {
	Source     string `gcfg:"src"`
	DestMovies string `gcfg:"dest-movies"`
	DestShows  string `gcfg:"dest-shows"`
}

type configServices struct {
	GuessItURL string `gcfg:"guessit"`
}

type configNames struct {
	DownloadsFolder string `gcfg:"downloads"`
	MoviesFolder    string `gcfg:"movies"`
	ShowsFolder     string `gcfg:"shows"`
}

type config struct {
	Paths    configPaths
	Services configServices
	Names    configNames
}

func readConfig(reader io.Reader) (*config, error) {
	cfg := new(config)
	if e := gcfg.ReadInto(cfg, reader); e != nil {
		return nil, e
	}
	cfg.setDefaultValues()
	return cfg, nil
}

func (c *config) setDefaultValues() {
	if (c.Names == configNames{}) {
		c.Names = configNames{
			DownloadsFolder: "Downloads",
			MoviesFolder:    "Movies",
			ShowsFolder:     "Shows",
		}
	}
}

func (c *config) parseGlobalSettings() GlobalSettings {
	return GlobalSettings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}
}

func (c *config) parseInputFoldersSettings() InputFoldersSettings {
	s := InputFoldersSettings{}
	if c.Paths.Source != "" {
		s.addPath(c.Names.DownloadsFolder, c.Paths.Source)
	}
	return s
}

func (c *config) parseOutputFoldersSettings() OutputFoldersSettings {
	s := OutputFoldersSettings{}
	if c.Paths.DestMovies != "" {
		s.addPath(c.Names.MoviesFolder, c.Paths.DestMovies)
	}
	if c.Paths.DestShows != "" {
		s.addPath(c.Names.ShowsFolder, c.Paths.DestShows)
	}
	return s
}

func (h *FoldersSettings) addPath(name, path string) {
	f := FolderInfo{Name: name, Path: utils.SplitPath(path)}
	h.Folders = append(h.Folders, f)
}
