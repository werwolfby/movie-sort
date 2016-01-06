package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	c := readConfig("config.ini")

	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.Handle("/api/settings", globalSettingsHandler{}).Methods("GET")
	r.Handle("/api/settings/input-folders", inputFoldersHandler{cfg: c}).Methods("GET")
	r.Handle("/api/settings/output-folders", outputFoldersHandler{cfg: c}).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":88", nil))
}
