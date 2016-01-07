package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGlobalSettings(t *testing.T) {
	s := globalSettings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}
	w := generateHandlerTest(t, "GET", (*globalSettingsHandler)(&s))

	assert.Equal(t, w.Code, 200)

	var actual globalSettings

	if e := json.NewDecoder(w.Body).Decode(&actual); e != nil {
		t.Errorf("%v", e)
	}

	assert.Equal(t, globalSettings{PathSeparator: fmt.Sprintf("%c", os.PathSeparator)}, actual)
}

func TestInputFoldersSettings(t *testing.T) {
	cfg := config{configPaths{Source: "D:\\Torrents\\Complete", DestMovies: "D:\\Video\\Movies", DestShows: "D:\\Video\\Shows"}}
	s := inputFoldersSettings{foldersSettings{cfg: &cfg}}
	s.init()

	w := generateHandlerTest(t, "GET", (*foldersHandler)(&s.foldersSettings))

	assert.Equal(t, w.Code, 200)

	var actual []folderInfo

	if e := json.NewDecoder(w.Body).Decode(&actual); e != nil {
		t.Errorf("%v", e)
	}

	assert.Equal(t, actual, []folderInfo{{"Downloads", []string{"D:", "Torrents", "Complete"}}})
}

func TestOutputFoldersSettings(t *testing.T) {
	cfg := config{configPaths{Source: "D:\\Torrents\\Complete", DestMovies: "D:\\Video\\Movies", DestShows: "D:\\Video\\Shows"}}
	s := outputFoldersSettings{foldersSettings{cfg: &cfg}}
	s.init()

	w := generateHandlerTest(t, "GET", (*foldersHandler)(&s.foldersSettings))

	assert.Equal(t, w.Code, 200)

	var actual []folderInfo

	if e := json.NewDecoder(w.Body).Decode(&actual); e != nil {
		t.Errorf("%v", e)
	}

	expected := []folderInfo{
		{"Movies", []string{"D:", "Video", "Movies"}},
		{"Shows", []string{"D:", "Video", "Shows"}}}

	assert.Equal(t, actual, expected)
}

func generateHandlerTest(t *testing.T, method string, handler http.Handler) *httptest.ResponseRecorder {
	req, e := http.NewRequest("GET", "", nil)
	if e != nil {
		t.Errorf("%v", e)
	}
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	return w
}
