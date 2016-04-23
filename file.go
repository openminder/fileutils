package fileutils

import (
	"os"
	"regexp"
	"strings"
)

// allowed image extensions
var imageExt = []string{"gif", "jpg", "jpeg", "png", "ico"}

// GetFileFromURL extracts a filename from the end of a URL
func GetFileFromURL(url string) string {
	r, err := regexp.Compile("[^/]+$")
	if err != nil {
		panic(err)
	}
	return r.FindString(url)
}

// FileIsImage checks if a filename has a valid image extension
func FileIsImage(filename string) bool {
	fileComponents := strings.Split(filename, ".")
	ext := fileComponents[len(fileComponents)-1]
	if stringInSlice(ext, imageExt) {
		return true
	}
	return false
}

// FileExists returns whether the given file or directory exists or not
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
