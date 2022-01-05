package fileprobe_test

import (
	"testing"

	"github.com/dollarsignteam/go-fileprobe"
)

func TestSetPath(t *testing.T) {
	expected := "/tmp/newPath"
	fp := fileprobe.NewHandler()
	fp.SetPath(expected)
	filePath := fp.GetPath()
	if filePath != expected {
		t.Errorf("expect %v but got %v", expected, filePath)
	}
}

func TestCreateFile(t *testing.T) {
	fp := fileprobe.NewHandler()
	filePath := fp.GetPath()
	expected := "/tmp/live"
	if filePath != expected {
		t.Errorf("expect %v but got %v", expected, filePath)
	}
	if fp.Exists() {
		t.Errorf("File exists without creating")
	}
	errCreate := fp.Create()
	if errCreate != nil {
		t.Error(errCreate)
	}
	if fp.Exists() != true {
		t.Errorf("File not found")
	}
	errRemove := fp.Remove()
	if errRemove != nil {
		t.Error(errRemove)
	}
}
