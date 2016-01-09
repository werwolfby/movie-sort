package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var splitPathTests = []struct {
	pathString string
	path       []string
}{
	{"D:\\Downloads\\Complete", []string{"D:", "Downloads", "Complete"}},
	{"D:\\Torrents\\Complete", []string{"D:", "Torrents", "Complete"}},
	// different slashes
	{"D:/Downloads/Complete", []string{"D:", "Downloads", "Complete"}},
	{"D:/Torrents/Complete", []string{"D:", "Torrents", "Complete"}},
	// mixed slashes
	{"D:\\Downloads/Complete", []string{"D:", "Downloads", "Complete"}},
	{"D:/Downloads\\Complete", []string{"D:", "Downloads", "Complete"}},
	{"D:\\Torrents/Complete", []string{"D:", "Torrents", "Complete"}},
	{"D:/Torrents\\Complete", []string{"D:", "Torrents", "Complete"}},
	// linux paths
	{"/mnt/media/Downloads/Complete", []string{"", "mnt", "media", "Downloads", "Complete"}},
	// UNC paths
	{"\\\\WORKSTATION\\media\\Downloads\\Complete", []string{"\\", "WORKSTATION", "media", "Downloads", "Complete"}},
}

func TestSplitPath1(t *testing.T) {
	for _, tt := range splitPathTests {
		path := splitPath(tt.pathString)
		assert.Equal(t, path, tt.path)
	}
}
