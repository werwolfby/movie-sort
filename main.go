package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	c := readConfig("config.ini")

	sh := settingsHandlers{}
	sh.init(c)

	r := mux.NewRouter()
	r.Handle("/api/settings", &sh.GlobalSettingsHandler).Methods("GET")
	r.Handle("/api/settings/input-folders", &sh.InputFoldersHandler).Methods("GET")
	r.Handle("/api/settings/output-folders", &sh.OutputFoldersHandler).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":88", nil))
}
