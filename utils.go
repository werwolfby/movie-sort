package main

import "strings"

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
	// path /mnt/path should be splitted to : ["", "mnt", "path"]
	// join of this slice will add prefix slash
	if len(s) >= 1 && s[0] == '/' {
		return append([]string{""}, result...)
	}

	return result
}
