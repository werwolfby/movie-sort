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

func TestAddPath(t *testing.T) {
	for _, tt := range addPathTests {
		f := foldersHandler{}
		for _, ap := range tt.AddPaths {
			f.addPath(ap.Name, ap.Path)
		}

		assert.Equal(t, f.folders, tt.Expected)
	}
}
