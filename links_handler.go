package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const linksAPI = "/api/links"

type linksHandlers struct {
	Links *links
}

type linkFunc func(oldname, newname string) error

func newLinksHandlers(l *links) *linksHandlers {
	return &linksHandlers{l}
}

func (l *linksHandlers) getLinksHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if len(path) < len(linksAPI) {
			w.WriteHeader(404)
			return
		}
		path = path[len(linksAPI):]
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		links := l.Links.getLinks(ioutil.ReadDir, os.SameFile, "", nil)
		json.NewEncoder(w).Encode(links)
	})
}

func (l *linksHandlers) getPutLinksHandler(link linkFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := vars["path"]
		splittedPath := splitPath(path)
		// there have to be at least two items in path: folder and filename
		if path == "" || len(splittedPath) < 2 {
			log.Printf("splittedPath is too small: %v", splittedPath)
			w.WriteHeader(400)
			return
		}
		inputFolder := splittedPath[0]
		inputFileInfo := fileInfo{inputFolder, splittedPath[1 : len(splittedPath)-1], splittedPath[len(splittedPath)-1], nil}
		inputFolderInfo := l.Links.Settings.InputFoldersSettings.find(inputFolder)
		if inputFolderInfo == nil {
			log.Printf("Can't find input folder: %s", inputFolder)
			w.WriteHeader(404)
			return
		}
		var linkFileInfo fileInfo
		if json.NewDecoder(r.Body).Decode(&linkFileInfo) != nil {
			log.Printf("Can't parse body")
			w.WriteHeader(400)
			return
		}
		outputFolderInfo := l.Links.Settings.OutputFoldersSettings.find(linkFileInfo.Folder)
		if outputFolderInfo == nil {
			log.Printf("Can't find output folder: %s", linkFileInfo.Folder)
			w.WriteHeader(404)
			return
		}
		srcPath := inputFileInfo.getFullName(*inputFolderInfo)
		dstPath := linkFileInfo.getFullName(*outputFolderInfo)

		dstDir := filepath.Dir(dstPath)

		if exists, _ := isDirExists(dstDir); !exists {
			if e := os.MkdirAll(dstDir, 0755); e != nil {
				w.WriteHeader(500)
				return
			}
		}

		log.Printf("Make link: %s -> %s", srcPath, dstPath)

		if e := link(srcPath, dstPath); e != nil {
			log.Printf("Make link: %s -> %s, Failed: %v", srcPath, dstPath, e)
			w.WriteHeader(500)
			return
		}

		resultLinkInfo := linkInfo{fileInfo: inputFileInfo, Links: []fileInfo{linkFileInfo}}

		json.NewEncoder(w).Encode(resultLinkInfo)
	})
}

func isDirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return stat != nil, err
}
