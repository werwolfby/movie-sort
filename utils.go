package main

import "strings"

const (
	pathSep = "\\/"
)

func splitPath(s string) []string {
	f := func(r rune) bool {
		for _, c := range pathSep {
			if r == c {
				return true
			}
		}
		return false
	}

	result := strings.FieldsFunc(s, f)

	// Prefix linux slash should not be removed
	// path /mnt/path should be splitted to : ["/mnt", "path"]
	// join of this slice will add prefix slash
	if len(s) >= 1 && s[0] == '/' {
		return append([]string{"/" + result[0]}, result[1:]...)
	}

	// Prefix UNC paths should not be removed
	// path \\WORKSTATION\path should be splitted to : ["\\WORKSTATION\path"]
	// join of this slice will return valid UNC path
	// this is required by filepath.Join method
	if strings.HasPrefix(s, "\\\\") && len(result) >= 2 {
		return append([]string{"\\\\" + result[0] + "\\" + result[1]}, result[2:]...)
	}

	return result
}
