// +build windows

package main

import "os"

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

func searchHardLinks(reader readDirFunc, sameFile sameFileFunc, inputFolders []folderInfo, outputFolders []folderInfo) []linkInfo {
	inputFiles := getAllFilesFrom(reader, inputFolders, extensions)
	outputFiles := getAllFilesFrom(reader, outputFolders, extensions)

	result := make([]linkInfo, len(inputFiles))

	for i, inputFile := range inputFiles {
		var links []fileInfo
		for _, outputFile := range outputFiles {
			if sameFile(inputFile.osFileInfo, outputFile.osFileInfo) {
				links = append(links, outputFile.myFileInfo)
			}
		}
		result[i] = linkInfo{fileInfo: inputFile.myFileInfo, Links: links}
	}

	return result
}
