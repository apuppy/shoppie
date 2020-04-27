package util

import (
	"path"
	"path/filepath"
	"strings"
)

// GetExt get extension from filename
func GetExt(filename string) string {
	return filepath.Ext(filename)
}

// GetFilename get filename from filename
func GetFilename(filename string) string {
	return strings.TrimSuffix(filename, path.Ext(filename))
}
