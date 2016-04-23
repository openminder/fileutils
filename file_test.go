package fileutils_test

import (
	"os"
	"testing"

	"github.com/openminder/fileutils"
)

func TestGetFileFromURL(t *testing.T) {
	v := fileutils.GetFileFromURL("http://localhost/image.png")
	if v != "image.png" {
		t.Error("Expected image.png, got", v)
	}

	v = fileutils.GetFileFromURL("http://localhost/file.csv")
	if v != "file.csv" {
		t.Error("Expected file.csv, got", v)
	}

	v = fileutils.GetFileFromURL("http://localhost/doc.pdf")
	if v != "doc.pdf" {
		t.Error("Expected doc.pdf, got", v)
	}

	v = fileutils.GetFileFromURL("http://localhost/")
	if v != "" {
		t.Error("Expected empty string, got", v)
	}
}

func TestFileIsImage(t *testing.T) {
	v := fileutils.FileIsImage("image.png")
	if !v {
		t.Error("Expected true, got", v)
	}
	v = fileutils.FileIsImage("image.ico")
	if !v {
		t.Error("Expected true, got", v)
	}
	v = fileutils.FileIsImage("image.jpg")
	if !v {
		t.Error("Expected true, got", v)
	}
	v = fileutils.FileIsImage("image.jpg")
	if !v {
		t.Error("Expected true, got", v)
	}
	v = fileutils.FileIsImage("image.jpeg")
	if !v {
		t.Error("Expected true, got", v)
	}
	v = fileutils.FileIsImage("image.gif")
	if !v {
		t.Error("Expected true, got", v)
	}
	v = fileutils.FileIsImage("image.pdf")
	if v {
		t.Error("Expected false, got", v)
	}
	v = fileutils.FileIsImage("image.docx")
	if v {
		t.Error("Expected false, got", v)
	}
}

func TestFileExists(t *testing.T) {
	var err error
	f, err := os.Create("testfile")
	if err != nil {
		panic(err)
	}
	f.Close()
	v, err := fileutils.FileExists("testfile")
	if !v {
		t.Error("Expected true, got false")
	}
	if err != nil {
		t.Error("Expected no error, got error")
	}
	err = os.Remove("testfile")
	if err != nil {
		panic(err)
	}
	v, err = fileutils.FileExists("testfile")
	if v {
		t.Error("Expected false, got true")
	}
	if err != nil {
		t.Error("Expected no error, got error")
	}
}

func TestGetExtensionFromFilename(t *testing.T) {
	ext := fileutils.GetExtensionFromFilename("file.doc")
	if ext != ".doc" {
		t.Error("Expected .doc, got ", ext)
	}
	ext = fileutils.GetExtensionFromFilename("file.xml")
	if ext != ".xml" {
		t.Error("Expected .xml, got ", ext)
	}
	ext = fileutils.GetExtensionFromFilename("file")
	if ext != "" {
		t.Error("Expected empty string, got ", ext)
	}
}
