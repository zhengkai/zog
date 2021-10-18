package zog

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// FileRotation ...
type FileRotation struct {
	mux     sync.Mutex
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

	name := fr.FnName(time.Now())

	fr.mux.Lock()
	defer fr.mux.Unlock()

	if name != fr.name {
		if fr.file != nil {
			fr.file.Close()
			fr.file = nil
		}

		dir := filepath.Dir(name)
		os.Mkdir(dir, fr.PermDir)

		var f *os.File
		f, err = os.OpenFile(name, baseFileMode|os.O_APPEND, fr.Perm)
		if err != nil {
			return
		}
		fr.file = f
		fr.name = name
	}

	return fr.file.Write(p)
}
