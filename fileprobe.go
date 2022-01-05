package fileprobe

import (
	"os"
)

const _path = "/tmp/live"

type handler struct {
	filePath string
}

// Create new probe handler
func NewHandler() *handler {
	return &handler{
		filePath: _path,
	}
}

// Set probe file path
func (h *handler) SetPath(path string) {
	h.filePath = path
}

// Get probe file path
func (h *handler) GetPath() string {
	return h.filePath
}

// Create probe file
func (h *handler) Create() error {
	_, err := os.Create(h.filePath)
	return err
}

// Remove probe file
func (h *handler) Remove() error {
	return os.Remove(h.filePath)
}

// Check if probe file exists
func (h *handler) Exists() bool {
	info, err := os.Stat(h.filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
