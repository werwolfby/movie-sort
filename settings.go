package main

import (
	"fmt"
	"os"
)

const (
	pathSep = "\\/"
)

type folderInfo struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type globalSettings struct {
	PathSeparator string `json:"pathSeparator"`
}

type foldersSettings struct {
	cfg     *config
	folders []folderInfo
}

type inputFoldersSettings struct {
	foldersSettings
}

type outputFoldersSettings struct {
	foldersSettings
}

type settings struct {
	GlobalSettings        globalSettings
	InputFoldersSettings  inputFoldersSettings
	OutputFoldersSettings outputFoldersSettings
}

func (sh *settings) init(cfg *config) {
	sh.InputFoldersSettings.cfg = cfg
	sh.OutputFoldersSettings.cfg = cfg

	sh.GlobalSettings.init()
	sh.InputFoldersSettings.init()
	sh.OutputFoldersSettings.init()
}

func newSettings(cfg *config) *settings {
	s := new(settings)
	s.init(cfg)
	return s
}

func (h *globalSettings) init() {
	h.PathSeparator = fmt.Sprintf("%c", os.PathSeparator)
}

func (h *foldersSettings) addPath(name, path string) {
	f := folderInfo{Name: name, Path: splitPath(path)}

	h.folders = append(h.folders, f)
}

func (h *inputFoldersSettings) init() {
	h.addPath("Downloads", h.cfg.Paths.Source)
}

func (h *outputFoldersSettings) init() {
	h.addPath("Movies", h.cfg.Paths.DestMovies)
	h.addPath("Shows", h.cfg.Paths.DestShows)
}
