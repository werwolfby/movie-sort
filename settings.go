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

type settings struct {
	PathSeparator string `json:"pathSeparator"`
}

type folderInfo struct {
	Name string   `json:"name"`
	Path []string `json:"path"`
}

type globalSettingsHandler struct {
}

type foldersHandler struct {
	cfg     *config
	folders []folderInfo
}

type inputFoldersHandler struct {
	foldersHandler
}

type outputFoldersHandler struct {
	foldersHandler
}

type settingsHandlers struct {
	GlobalSettingsHandler globalSettingsHandler
	InputFoldersHandler   inputFoldersHandler
	OutputFoldersHandler  outputFoldersHandler
}

func (sh *settingsHandlers) init(cfg *config) {
	sh.InputFoldersHandler.cfg = cfg
	sh.OutputFoldersHandler.cfg = cfg

	sh.InputFoldersHandler.init()
	sh.OutputFoldersHandler.init()
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

	return strings.FieldsFunc(s, f)
}

func (h *foldersHandler) addPath(name, path string) {
	f := folderInfo{Name: name, Path: splitPath(path)}

	h.folders = append(h.folders, f)
}

func (h *inputFoldersHandler) init() {
	h.addPath("Downloads", h.cfg.Paths.Source)
}

func (h *outputFoldersHandler) init() {
	h.addPath("Movies", h.cfg.Paths.DestMovies)
	h.addPath("Shows", h.cfg.Paths.DestShows)
}

func (*globalSettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := settings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}
	json.NewEncoder(w).Encode(s)
}

func (h *foldersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.folders)
}
