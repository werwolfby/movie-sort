package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	c, e := readConfig("config.ini")
	if e != nil {
		log.Fatal(e)
	}

	s := newSettings(c)
	sh := (*settingsHandler)(s)

	r := mux.NewRouter()

	r.Handle("/api/settings", sh.getGlobalSettingsHandler()).Methods("GET")
	r.Handle("/api/settings/input-folders", sh.getInputFolderSettingsHandler()).Methods("GET")
	r.Handle("/api/settings/output-folders", sh.getOutputFolderSettingsHandler()).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	http.Handle("/", r)
	log.Println("Starting server on http://localhost:88")
	log.Fatal(http.ListenAndServe(":88", nil))
}
