package zog

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const baseFileMode = os.O_WRONLY | os.O_CREATE

// default permission when create file / dir
var (
	DefaultPermFile os.FileMode = 0644
	DefaultPermDir  os.FileMode = 0755
)

// FileRotation ...
type FileRotation struct {
	mux      sync.Mutex
	FnName   func(time.Time) string
	name     string
	file     *os.File
	PermFile os.FileMode
	PermDir  os.FileMode
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
		PermFile: DefaultPermFile,
		PermDir:  DefaultPermDir,
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
		f, err = os.OpenFile(name, baseFileMode|os.O_APPEND, fr.PermFile)
		if err != nil {
			return
		}
		fr.file = f
		fr.name = name
	}

	return fr.file.Write(p)
}

// NewFile utils for create log file, will create dir automatically
func NewFile(name string, isAppend bool) (f *os.File, err error) {

	err = os.MkdirAll(filepath.Dir(name), DefaultPermDir)
	if err != nil {
		return
	}

	flag := baseFileMode
	if isAppend {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}
	return os.OpenFile(name, flag, DefaultPermFile)
}
