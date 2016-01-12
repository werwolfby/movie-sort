// +build windows

package main

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
