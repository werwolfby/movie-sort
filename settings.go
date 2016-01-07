package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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

type settingsHandler settings

type globalSettingsHandler globalSettings

type foldersHandler foldersSettings

func (sh *settings) init(cfg *config) {
	sh.InputFoldersSettings.cfg = cfg
	sh.OutputFoldersSettings.cfg = cfg

	sh.GlobalSettings.init()
	sh.InputFoldersSettings.init()
	sh.OutputFoldersSettings.init()
}

func splitPath(s string) []string {
	f := func(r rune) bool {
		for _, c := range pathSep {
			if r == c {
				return true
			}
		}
		return false
	}

	result := strings.FieldsFunc(s, f)

	// Prefix linux slash should not be removed
	// path /mnt/path should be splitted to : ["", "mnt", "path"]
	// join of this slice will add prefix slash
	if len(s) >= 1 && s[0] == '/' {
		return append([]string{""}, result...)
	}

	return result
}

func newSettings(cfg *config) *settings {
	s := new(settings)
	s.init(cfg)
	return s
}

func (s *settingsHandler) getGlobalSettingsHandler() http.Handler {
	return (*globalSettingsHandler)(&s.GlobalSettings)
}

func (s *settingsHandler) getInputFolderSettingsHandler() http.Handler {
	return (*foldersHandler)(&s.InputFoldersSettings.foldersSettings)
}

func (s *settingsHandler) getOutputFolderSettingsHandler() http.Handler {
	return (*foldersHandler)(&s.OutputFoldersSettings.foldersSettings)
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

func (h *globalSettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h)
}

func (h *foldersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.folders)
}
