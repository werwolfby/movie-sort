package server

import (
	"github.com/werwolfby/movie-sort/settings"
	"net/http"
)

type SettingsHandlers struct {
	settings *settings.ApplicationSettings
}

func NewSettingsHandlers(s *settings.ApplicationSettings) *SettingsHandlers {
	return &SettingsHandlers{settings: s}
}

func (h SettingsHandlers) GetGlobalSettingsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeOk(w, h.settings.Global)
	})
}

func (h SettingsHandlers) GetInputFoldersSettingsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeOk(w, h.settings.InputFolders)
	})
}

func (h SettingsHandlers) GetOutputFoldersSettingsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeOk(w, h.settings.OutputFolders)
	})
}
