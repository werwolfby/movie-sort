package settings

import (
	"io"
	"os"
)

type FolderNames struct {
	DownloadsName string
	MoviesName    string
	ShowsName     string
}

type FolderInfo struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type GlobalSettings struct {
	PathSeparator string `json:"pathSeparator"`
}

type FoldersSettings struct {
	Folders []FolderInfo
}

type InputFoldersSettings struct {
	FoldersSettings
}

type OutputFoldersSettings struct {
	FoldersSettings
}

type Services struct {
	GuessItURL string
}

type Settings struct {
	Global        GlobalSettings        `json:"global"`
	InputFolders  InputFoldersSettings  `json:"input-folders"`
	OutputFolders OutputFoldersSettings `json:"output-folders"`
}

type ApplicationSettings struct {
	Settings
	FolderNames FolderNames
	Services    Services
}

func ReadSettingsFromFile(filename string) (*ApplicationSettings, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadSettings(f)
}

func ReadSettings(reader io.Reader) (*ApplicationSettings, error) {
	cfg, err := readConfig(reader)
	if err != nil {
		return nil, err
	}
	s := ApplicationSettings{
		Settings: Settings{
			Global:        cfg.parseGlobalSettings(),
			InputFolders:  cfg.parseInputFoldersSettings(),
			OutputFolders: cfg.parseOutputFoldersSettings(),
		},
		FolderNames: FolderNames{
			DownloadsName: cfg.Names.DownloadsFolder,
			MoviesName:    cfg.Names.MoviesFolder,
			ShowsName:     cfg.Names.ShowsFolder,
		},
		Services: Services{
			GuessItURL: cfg.Services.GuessItURL,
		},
	}
	return &s, nil
}
