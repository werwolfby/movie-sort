package main

import (
	"encoding/json"
	"net/http"
)

type settingsHandler struct {
	settings *settings
}

func newSettingsHandler(s *settings) *settingsHandler {
	sh := new(settingsHandler)
	sh.settings = s
	return sh
}

func (s *settingsHandler) getGlobalSettingsHandler() http.Handler {
	return makeJSONHandler(func() interface{} { return s.settings.GlobalSettings })
}

func (s *settingsHandler) getInputFoldersSettingsHandler() http.Handler {
	return makeJSONHandler(func() interface{} { return s.settings.InputFoldersSettings.folders })
}

func (s *settingsHandler) getOutputFoldersSettingsHandler() http.Handler {
	return makeJSONHandler(func() interface{} { return s.settings.OutputFoldersSettings.folders })
}

func makeJSONHandler(f func() interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(f())
	})
}
