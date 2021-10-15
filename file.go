package zog

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileRotation ...
type FileRotation struct {
	FnName  func(time.Time) string
	name    string
	file    *os.File
	Perm    os.FileMode
	PermDir os.FileMode
}

// NewFileRotation ...
func NewFileRotation(dir, format string) *FileRotation {
	if !strings.HasSuffix(dir, `/`) {
		dir += `/`
	}
	return &FileRotation{
		FnName: func(t time.Time) string {
			return dir + t.Format(format)
		},
		Perm:    0644,
		PermDir: 0755,
	}
}

// Write ...
func (fr *FileRotation) Write(p []byte) (n int, err error) {
	return fr.TimeWrite(p, time.Now())
}

// TimeWrite ...
func (fr *FileRotation) TimeWrite(p []byte, t time.Time) (n int, err error) {

	name := fr.FnName(t)
	if name != fr.name {
		if fr.file != nil {
			fr.file.Close()
			fr.file = nil
		}

		dir := filepath.Dir(name)
		os.Mkdir(dir, fr.PermDir)

		var f *os.File
		f, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE, fr.Perm)
		if err != nil {
			return
		}
		fr.file = f
		fr.name = name
	}

	return fr.file.Write(p)
}
