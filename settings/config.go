package settings

import (
	"fmt"
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
	if (cfg.Names == configNames{}) {
		cfg.Names = configNames{
			DownloadsFolder: "Downloads",
			MoviesFolder:    "Movies",
			ShowsFolder:     "Shows",
		}
	}
	return cfg, nil
}

func (c *config) parseGlobalSettings() GlobalSettings {
	return GlobalSettings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}
}

func (c *config) parseInputFoldersSettings() InputFoldersSettings {
	s := InputFoldersSettings{}
	if c.Paths.Source != "" {
		s.addPath(c.Names.DownloadsFolder, c.Paths.Source, FolderMetaDownloads)
	}
	return s
}

func (c *config) parseOutputFoldersSettings() OutputFoldersSettings {
	s := OutputFoldersSettings{}
	if c.Paths.DestMovies != "" {
		s.addPath(c.Names.MoviesFolder, c.Paths.DestMovies, FolderMetaMovies)
	}
	if c.Paths.DestShows != "" {
		s.addPath(c.Names.ShowsFolder, c.Paths.DestShows, FolderMetaShows)
	}
	return s
}
