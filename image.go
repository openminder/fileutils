package fileutils

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// GenerateThumbnail generates a thumbanil of a given image path
func GenerateThumbnail(sourceFilePath, targetFilePath string, width uint, height uint) {
	reader, err := os.Open(sourceFilePath)
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	newImage := resize.Thumbnail(width, height, m, resize.Lanczos3)
	extension := GetExtensionFromFilename(GetFileFromURL(sourceFilePath))
	buf := new(bytes.Buffer)
	switch extension {
	case ".png":
		png.Encode(buf, newImage)
	case ".jpg":
		jpeg.Encode(buf, newImage, nil)
	case ".jpeg":
		jpeg.Encode(buf, newImage, nil)
	case ".gif":
		gif.Encode(buf, newImage, nil)
	default:
		err = errors.New("image: unknown format: " + extension)
	}
	if err != nil {
		log.Fatal(err)
	}
	imgAsByte := buf.Bytes()
	filename := GetFileFromURL(sourceFilePath)
	exists, err := FileExists(targetFilePath)
	if exists != true {
		os.MkdirAll(targetFilePath, 0777)
	}
	if err != nil {
		log.Fatal(err)
	}
	_, err = SaveToDisc(targetFilePath+filename, imgAsByte)
	if err != nil {
		log.Fatal(err)
	}
}
