package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type addPathInfo struct {
	Name string
	Path string
}

var addPathTests = []struct {
	AddPaths []addPathInfo
	Expected []folderInfo
}{
	{[]addPathInfo{{"Test1", "D:\\Downloads\\Test"}}, []folderInfo{{"Test1", []string{"D:", "Downloads", "Test"}}}},
	{[]addPathInfo{{"Downloads", "D:\\Downloads\\Test"}}, []folderInfo{{"Downloads", []string{"D:", "Downloads", "Test"}}}},
	{[]addPathInfo{
		{"Downloads1", "D:\\Downloads\\Test1"},
		{"Downloads2", "D:\\Downloads\\Test2"}},
		[]folderInfo{
			{"Downloads1", []string{"D:", "Downloads", "Test1"}},
			{"Downloads2", []string{"D:", "Downloads", "Test2"}}}},
}

var inputFoldersInitTests = []struct {
	Config   config
	Expected []folderInfo
}{
	{config{Paths: configPaths{Source: "D:\\Torrents\\Complete", DestMovies: "D:\\Video\\Movies", DestShows: "D:\\Video\\Shows"}},
		[]folderInfo{{"Downloads", []string{"D:", "Torrents", "Complete"}}}},
	{config{Paths: configPaths{Source: "/mnt/media/Torrents/Complete", DestMovies: "/mnt/media/Video/Movies", DestShows: "/mnt/media/Video/Shows"}},
		[]folderInfo{{"Downloads", []string{"/mnt", "media", "Torrents", "Complete"}}}},
}

var outputFoldersInitTests = []struct {
	Config   config
	Expected []folderInfo
}{
	{config{Paths: configPaths{Source: "D:\\Torrents\\Complete", DestMovies: "D:\\Video\\Movies", DestShows: "D:\\Video\\Shows"}},
		[]folderInfo{
			{"Movies", []string{"D:", "Video", "Movies"}},
			{"Shows", []string{"D:", "Video", "Shows"}}}},
	{config{Paths: configPaths{Source: "/mnt/media/Torrents/Complete", DestMovies: "/mnt/media/Video/Movies", DestShows: "/mnt/media/Video/Shows"}},
		[]folderInfo{
			{"Movies", []string{"/mnt", "media", "Video", "Movies"}},
			{"Shows", []string{"/mnt", "media", "Video", "Shows"}}}},
}

func TestAddPath(t *testing.T) {
	for _, tt := range addPathTests {
		f := foldersSettings{}
		for _, ap := range tt.AddPaths {
			f.addPath(ap.Name, ap.Path)
		}

		assert.Equal(t, f.folders, tt.Expected)
	}
}

func TestInputFoldersInit(t *testing.T) {
	for _, tt := range inputFoldersInitTests {
		f := inputFoldersSettings{foldersSettings{cfg: &tt.Config}}
		f.init()

		assert.Equal(t, f.folders, tt.Expected)
	}
}

func TestOutputFoldersInit(t *testing.T) {
	for _, tt := range outputFoldersInitTests {
		f := outputFoldersSettings{foldersSettings{cfg: &tt.Config}}
		f.init()

		assert.Equal(t, f.folders, tt.Expected)
	}
}
