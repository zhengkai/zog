package zog

import (
	"fmt"
)

// Level iota
const (
	LevelDefault = Level(iota)
	LevelWrite
	LevelPrint
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
	CTrace   *Config
	CDebug   *Config
	CInfo    *Config
	CError   *Config
	CWarn    *Config
	CFatal   *Config
}

// NewSimple ...
func NewSimple(file string) (i *Input, err error) {

	f, err := OutputFile(file, false)
	if err != nil {
		return
	}

	cfg := NewConfig()
	cfg.AddOutput(f)

	i = &Input{
		CDefault: cfg,
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

// SetDirPrefix ...
func (i *Input) SetDirPrefix(d string) {
	for _, c := range []*Config{i.CDefault, i.CWrite, i.CPrint, i.CDebug, i.CInfo, i.CError, i.CWarn, i.CFatal} {
		if c != nil {
			c.SetDirPrefix(d)
		}
	}
}
