package zog

import (
	"path/filepath"
	"runtime"
)

// GetSourceFileDir ...
func GetSourceFileDir() string {
	_, file, _, _ := runtime.Caller(1)
	dir, _ := filepath.Abs(filepath.Dir(file))
	return dir
}
