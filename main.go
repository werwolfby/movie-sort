package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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

func main() {
	c := readConfig("config.ini")

	sh := settingsHandlers{}
	sh.init(c)

	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.Handle("/api/settings", &sh.GlobalSettingsHandler).Methods("GET")
	r.Handle("/api/settings/input-folders", &sh.InputFoldersHandler).Methods("GET")
	r.Handle("/api/settings/output-folders", &sh.OutputFoldersHandler).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":88", nil))
}
