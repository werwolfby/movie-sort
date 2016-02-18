// +build windows

package links

func (l *links) searchLinks(inputFiles []searchFileInfo, outputFiles []searchFileInfo) []LinkInfo {
	result := make([]LinkInfo, len(inputFiles))
	for i, isfi := range inputFiles {
		var links []FileInfo
		for _, osfi := range outputFiles {
			// Check for size before SameFile is just a small optimization
			if isfi.OsFileInfo.Size() == osfi.OsFileInfo.Size() && l.Reader.SameFile(isfi.OsFileInfo, osfi.OsFileInfo) {
				links = append(links, osfi.FileInfo)
			}
		}
		result[i] = LinkInfo{FileInfo: isfi.FileInfo, Links: links}
	}
	return result
}
