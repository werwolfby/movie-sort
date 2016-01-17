package guessit

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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

	guessit := NewGuessItService(client.URL)
	result, err := guessit.GuessIt("filename.avi")

	assert.Nil(t, err)
	assert.Equal(t, "episode", result.Type)
	assert.Equal(t, "Test", result.Title)
	assert.Equal(t, 3, result.Season)
}

func handleGuess(resp interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(resp)
	})
}
