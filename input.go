package zog

import (
	"fmt"
	"io"
	"os"
)

// Level iota
const (
	LevelDefault = Level(iota)
	LevelDebug
	LevelInfo
	LevelError
	LevelWarn
	LevelFatal
)

// Level ...
type Level uint8

// Input ...
type Input struct {
	CDefault *Config
	CWrite   *Config
	CPrint   *Config
	CDebug   *Config
	CInfo    *Config
	CError   *Config
	CWarn    *Config
	CFatal   *Config

	DirBase string
}

// NewSimple ...
func NewSimple(name string) (i *Input, err error) {

	f, err := OutputFile(name, false)
	if err != nil {
		return
	}

	i = &Input{
		CDefault: &Config{
			TimeFormat: `2006-01-02 15:04:05 `,
			Caller:     CallerShorter,
			Output: &Output{
				Dest: []io.Writer{
					os.Stdout,
					f,
				},
			},
		},
	}

	return
}

// Print ...
func (i *Input) Print(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CPrint, s)
}

// Println ...
func (i *Input) Println(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CPrint, s)
}

// Writer for io.Writer
func (i *Input) Write(p []byte) (n int, err error) {
	c := i.CWrite
	if c == nil {
		c = i.CDefault
	}
	c.writeAB(p)
	n = len(p)
	return
}

func (i *Input) write(c *Config, p string) {
	if c == nil {
		c = i.CDefault
	}
	c.write(p)
}
