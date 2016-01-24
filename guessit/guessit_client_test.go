package guessit

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/werwolfby/movie-sort/links"
	"github.com/werwolfby/movie-sort/settings"
	"net/http"
	"net/http/httptest"
	"testing"
)

type guessitResponse struct {
	Type           string `json:"type"`
	Title          string `json:"title"`
	Season         int    `json:"season"`
	SomeotherField string `json:"someother_field"`
}

func TestGuessitClient(t *testing.T) {
	resp := guessitResponse{Type: "episode", Title: "Test", Season: 3, SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	guessit := guessItService{serviceURL: client.URL}
	result, err := guessit.guessIt("filename.avi")

	assert.Nil(t, err)
	assert.Equal(t, "episode", result.Type)
	assert.Equal(t, "Test", result.Title)
	assert.Equal(t, 3, result.Season)
}

func TestGuessitClientErr(t *testing.T) {
	var client = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		err := struct {
			Code  int
			Error string
		}{500, "Internal Server Error"}
		json.NewEncoder(w).Encode(err)
	}))
	defer client.Close()

	guessit := guessItService{serviceURL: client.URL}
	result, err := guessit.guessIt("filename.avi")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestGuessitClientErr2(t *testing.T) {
	var client = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	guessit := guessItService{serviceURL: client.URL}
	client.Close()
	result, err := guessit.guessIt("filename.avi")

	assert.NotNil(t, err)
	assert.Nil(t, result)
}

type mockLinks struct {
	GetShowsFolderResult  settings.FolderInfo
	GetShowsFolderError   bool
	GetMoviesFolderResult settings.FolderInfo
	GetMoviesFolderError  bool
	GetShowSeasonFileInfo links.FileInfo
	GetShowSeasonSeason   int
}

func (m mockLinks) GetShowsFolder() (settings.FolderInfo, bool) {
	return m.GetShowsFolderResult, m.GetShowsFolderError
}

func (m mockLinks) GetMoviesFolder() (settings.FolderInfo, bool) {
	return m.GetMoviesFolderResult, m.GetMoviesFolderError
}

func (m mockLinks) GetShowSeason(name string, season int) (links.FileInfo, int) {
	return m.GetShowSeasonFileInfo, m.GetShowSeasonSeason
}

func TestGuessExistingSeasonLink(t *testing.T) {
	resp := guessitResponse{Type: "episode", Title: "Test", Season: 3, SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)

	mock.GetShowSeasonFileInfo = links.FileInfo{Folder: "Shows", Path: []string{"Test"}, Name: "Season 3"}
	mock.GetShowSeasonSeason = 3

	fi, err := guessit.GuessLink("Test.avi")

	assert.NotNil(t, fi)
	assert.Nil(t, err)

	assert.Equal(t, "Shows", fi.Folder)
	assert.Equal(t, []string{"Test", "Season 3"}, fi.Path)
	assert.Equal(t, "Test.avi", fi.Name)
}

func TestGuessExistingShowLink(t *testing.T) {
	resp := guessitResponse{Type: "episode", Title: "Test", Season: 3, SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)

	mock.GetShowSeasonFileInfo = links.FileInfo{Folder: "Shows", Path: []string{}, Name: "Test"}
	mock.GetShowSeasonSeason = links.SeasonNotFound

	fi, err := guessit.GuessLink("Test.avi")

	assert.NotNil(t, fi)
	assert.Nil(t, err)

	assert.Equal(t, "Shows", fi.Folder)
	assert.Equal(t, []string{"Test", "Season 3"}, fi.Path)
	assert.Equal(t, "Test.avi", fi.Name)
}

func TestGuessNewShowLink(t *testing.T) {
	resp := guessitResponse{Type: "episode", Title: "Test", Season: 3, SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)

	mock.GetShowSeasonSeason = links.ShowNotFound

	mock.GetShowsFolderResult = settings.FolderInfo{Name: "Shows", Path: []string{}, Meta: settings.FolderMetaShows}
	mock.GetShowsFolderError = true

	fi, err := guessit.GuessLink("Test.avi")

	assert.NotNil(t, fi)
	assert.Nil(t, err)

	assert.Equal(t, "Shows", fi.Folder)
	assert.Equal(t, []string{"Test", "Season 3"}, fi.Path)
	assert.Equal(t, "Test.avi", fi.Name)
}

func TestGuessNewShowLinkErr(t *testing.T) {
	resp := guessitResponse{Type: "episode", Title: "Test", Season: 3, SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)

	mock.GetShowSeasonSeason = links.ShowNotFound

	mock.GetShowsFolderError = false

	fi, err := guessit.GuessLink("Test.avi")

	assert.Nil(t, fi)
	assert.NotNil(t, err)
}

func TestGuessNewMovieLink(t *testing.T) {
	resp := guessitResponse{Type: "movie", Title: "Test", SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)

	mock.GetShowSeasonSeason = links.ShowNotFound

	mock.GetShowsFolderResult = settings.FolderInfo{Name: "Shows", Path: []string{}, Meta: settings.FolderMetaShows}
	mock.GetShowsFolderError = true

	mock.GetMoviesFolderResult = settings.FolderInfo{Name: "Movies", Path: []string{}, Meta: settings.FolderMetaMovies}
	mock.GetMoviesFolderError = true

	fi, err := guessit.GuessLink("Test.avi")

	assert.NotNil(t, fi)
	assert.Nil(t, err)

	assert.Equal(t, "Movies", fi.Folder)
	assert.Equal(t, []string{}, fi.Path)
	assert.Equal(t, "Test.avi", fi.Name)
}

func TestGuessNewMovieLinkErr(t *testing.T) {
	resp := guessitResponse{Type: "movie", Title: "Test", SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))
	defer client.Close()

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)

	mock.GetShowSeasonSeason = links.ShowNotFound

	mock.GetShowsFolderError = false

	mock.GetMoviesFolderError = false

	fi, err := guessit.GuessLink("Test.avi")

	assert.Nil(t, fi)
	assert.NotNil(t, err)
}

func TestGuessLinkErr(t *testing.T) {
	resp := guessitResponse{Type: "movie", Title: "Test", SomeotherField: "AsdF"}
	var client = httptest.NewServer(handleGuess(resp))

	mock := &mockLinks{}

	guessit := NewGuessItService(client.URL, mock)
	client.Close()

	fi, err := guessit.GuessLink("Test.avi")

	assert.Nil(t, fi)
	assert.NotNil(t, err)
}

func handleGuess(resp interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(resp)
	})
}
