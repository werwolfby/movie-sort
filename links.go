package main

import "os"

type linkInfo struct {
	fileInfo
	Links []fileInfo `json:"links"`
}

type links struct {
	Settings *settings
}

func newLinks(s *settings) *links {
	return &links{s}
}

func (l *links) getLinks(reader readDirFunc, sameFile sameFileFunc, folder string, path []string) []linkInfo {
	return searchHardLinks(reader, sameFile,
		l.Settings.InputFoldersSettings.folders,
		l.Settings.OutputFoldersSettings.folders)
}

var extensions = []string{"mkv", "avi", "mp4"}

type searchFileInfo struct {
	osFileInfo   os.FileInfo
	myFolderInfo *folderInfo
	myFileInfo   fileInfo
}

func getAllFilesFrom(reader readDirFunc, folders []folderInfo, extensions []string) []searchFileInfo {
	var files = []searchFileInfo{}
	for _, folder := range folders {
		folderFiles, e := getAllFiles(reader, folder.Name, folder.Path, []string{}, extensions)
		if e == nil {
			for _, file := range folderFiles {
				files = append(files, searchFileInfo{file.OsFileInfo, &folder, file})
			}
		}
	}
	return files
}
