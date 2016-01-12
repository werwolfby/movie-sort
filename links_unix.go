// +build linux

package main

import (
	"syscall"
)

func searchHardLinks(reader readDirFunc, sameFile sameFileFunc, inputFolders []folderInfo, outputFolders []folderInfo) []linkInfo {
	createMap := func(files []searchFileInfo) map[uint64][]fileInfo {
		result := map[uint64][]fileInfo{}
		for _, fi := range files {
			if sys := fi.osFileInfo.Sys(); sys != nil {
				if stat, ok := sys.(*syscall.Stat_t); ok {
					result[uint64(stat.Ino)] = append(result[uint64(stat.Ino)], fi.myFileInfo)
				}
			}
		}
		return result
	}

	inputFiles := getAllFilesFrom(reader, inputFolders, extensions)
	outputFiles := getAllFilesFrom(reader, outputFolders, extensions)

	sourceInos := createMap(inputFiles)
	destInos := createMap(outputFiles)

	result := []linkInfo{}

	for ino, files := range sourceInos {
		links := destInos[ino]
		for _, f := range files {
			infos := linkInfo{fileInfo: f, Links: links}
			result = append(result, infos)
		}
	}

	return result
}
