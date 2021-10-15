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
	CDefault *config
	CPrint   *config
	CDebug   *config
	CInfo    *config
	CError   *config
	CWarn    *config
	CFatal   *config

	DirBase string
}

// NewSimple ...
func NewSimple(name string) (i *Input, err error) {

	f, err := OutputFile(name, false)
	if err != nil {
		return
	}

	i = &Input{
		CDefault: &config{
			o: &output{
				dst: []io.Writer{
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
	i.write(i.CPrint, []byte(s))
}

func (i *Input) write(c *config, p []byte) {
	if c == nil {
		c = i.CDefault
	}
	c.write(p)
}
