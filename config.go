package zog

import (
	"bytes"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const defaultTimeFormat = `2006-01-02 15:04:05 `

// Config ...
type Config struct {
	Caller          CallerType
	Output          []io.Writer
	IgnoreOutputErr bool
	LinePrefix      string // beginning of the line
	MsgPrefix       string // before the message
	Color           string
	TimeFormat      string
	dir             string
	dirLen          int
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Caller:     CallerLong,
		TimeFormat: defaultTimeFormat,
		Output: []io.Writer{
			os.Stdout,
		},
	}
}

// AddOutput ...
func (c *Config) AddOutput(w io.Writer) {
	c.Output = append(c.Output, w)
}

// Clone ...
func (c *Config) Clone() *Config {
	x := *c
	x.Output = nil
	for _, v := range c.Output {
		x.AddOutput(v)
	}
	return &x
}

func (c *Config) writePrepare() (buf bytes.Buffer) {

	if c.Color != `` {
		buf.WriteString("\x1b[")
		buf.WriteString(c.Color)
		buf.WriteRune('m')
	}

	if c.LinePrefix != `` {
		buf.WriteString(c.LinePrefix)
	}

	if c.TimeFormat != `` {
		buf.WriteString(time.Now().Format(c.TimeFormat))
	}

	if c.Caller != CallerNone {
		_, file, line, ok := runtime.Caller(4)
		if ok {
			file = strings.TrimSuffix(file, hideExt)
			if c.dirLen > 0 && strings.HasPrefix(file, c.dir) {
				file = file[c.dirLen:]
			}
			buf.WriteString(file)
			buf.WriteRune(':')
			i := strconv.Itoa(line)
			buf.WriteString(i)
		} else {
			buf.WriteString(`unknown:0`)
		}
		buf.WriteRune(' ')
	}

	if c.MsgPrefix != `` {
		buf.WriteString(c.MsgPrefix)
	}

	return
}

func (c *Config) writeAB(msg []byte) {

	buf := c.writePrepare()

	l := len(msg)
	if l == 0 {
		// do nothing
	} else if msg[l-1] == '\n' {
		buf.Write(msg[:l-1])
	} else {
		buf.Write(msg)
	}
	if c.Color != `` {
		buf.WriteString("\x1b[0m")
	}
	buf.WriteRune('\n')

	buf.WriteTo(c)
}

func (c *Config) write(msg string) {

	buf := c.writePrepare()

	if msg != `` {
		if strings.HasSuffix(msg, "\n") {
			buf.WriteString(msg[:len(msg)-1])
		} else {
			buf.WriteString(msg)
		}
	}

	if c.Color != `` {
		buf.WriteString("\x1b[0m")
	}

	buf.WriteRune('\n')

	buf.WriteTo(c)
}

func (c *Config) Write(p []byte) (n int, err error) {

	for _, o := range c.Output {
		n, err = o.Write(p)
		if err != nil && !c.IgnoreOutputErr {
			return
		}
	}
	return
}

// SetDirPrefix dir prefix in caller filename with be hidden
func (c *Config) SetDirPrefix(d string) {

	if len(d) == 0 {
		c.dir = ``
		c.dirLen = 0
		return
	}

	if !strings.HasSuffix(d, `/`) {
		d += `/`
	}
	c.dir = d
	c.dirLen = len(d)
}
