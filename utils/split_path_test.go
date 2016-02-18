package utils_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/werwolfby/movie-sort/utils"
	"testing"
)

var splitPathTests = map[string][]string{
	"D:\\Downloads\\Complete":                     []string{"D:", "Downloads", "Complete"},
	"D:\\Torrents\\Complete":                      []string{"D:", "Torrents", "Complete"},
	"D:/Downloads/Complete":                       []string{"D:", "Downloads", "Complete"}, // different slashes
	"D:/Torrents/Complete":                        []string{"D:", "Torrents", "Complete"},
	"D:\\Downloads/Complete":                      []string{"D:", "Downloads", "Complete"}, // mixed slashes
	"D:/Downloads\\Complete":                      []string{"D:", "Downloads", "Complete"},
	"D:\\Torrents/Complete":                       []string{"D:", "Torrents", "Complete"},
	"D:/Torrents\\Complete":                       []string{"D:", "Torrents", "Complete"},
	"/mnt/media/Downloads/Complete":               []string{"/mnt", "media", "Downloads", "Complete"},          // linux paths
	"\\\\WORKSTATION\\media\\Downloads\\Complete": []string{"\\\\WORKSTATION\\media", "Downloads", "Complete"}, // UNC paths
	"\\\\WORKSTATION\\media":                      []string{"\\\\WORKSTATION\\media"},                          // single path for corner cases
	"/mnt": []string{"/mnt"},
}

func TestSplitPath1(t *testing.T) {
	for k, v := range splitPathTests {
		path := utils.SplitPath(k)
		assert.Equal(t, v, path)
	}
}
