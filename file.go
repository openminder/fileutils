package fileutils

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
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
	newURL := r.FindString(url)
	newURL = strings.Split(newURL, "?")[0]
	return newURL
}

// GetExtensionFromFilename extracts extension from filename
func GetExtensionFromFilename(filename string) string {
	// check if file has extension
	if !strings.Contains(filename, ".") {
		return ""
	}
	r, err := regexp.Compile(".[0-9a-z]+$")
	if err != nil {
		panic(err)
	}
	return r.FindString(filename)
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

// stringInSlice checks if a string is in slice of strings
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// DownloadFile saves a file on the file system and returns the files size or an error
func DownloadFile(source string, targetPath string, filename string) (int64, error) {
	if source == "" {
		return 0, nil
	}
	if targetPath == "" {
		return 0, nil
	}
	if filename == "" {
		filename = GetFileFromURL(source)
	}
	if string(targetPath[len(targetPath)-1]) != "/" {
		targetPath = targetPath + "/"
	}
	exists, err := FileExists(targetPath)
	if exists != true {
		os.MkdirAll(targetPath, 0777)
	}
	if err != nil {
		return 0, err
	}
	resp, err := http.Get(source)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	size, err := SaveToDisc(targetPath+filename, body)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// SaveToDisc saves a file to the filesystem
func SaveToDisc(filePath string, fileContent []byte) (int64, error) {
	out, err := os.Create(filePath)
	defer out.Close()
	if err != nil {
		return 0, err
	}
	r := bytes.NewReader(fileContent)
	size, err := io.Copy(out, r)
	if err != nil {
		return 0, err
	}
	return size, nil
}
