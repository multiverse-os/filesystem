package memfs

import "strings"

func PopPath(path string) (string, string) {
	if len(path) == 0 {
		return "", "" // 1
	}
	if path == "/" {
		return "/", "" // 2
	}

	x := strings.Index(path, "/")
	if x == -1 {
		return path, "" // 6
	} else if x == 0 {
		return "/", strings.TrimLeft(path, "/") // 3
	}
	return path[:x], path[x+1:] // 4, 5
}
