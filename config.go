package zog

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Error ...
var (
	ErrNotInited = errors.New(`not initialized`)
	ErrNoOutput  = errors.New(`no output`)
	ErrDisabled  = errors.New(`disabled`)
)

// Config ...
type Config struct {
	Caller          CallerType
	Output          []io.Writer
	Enable          bool
	IgnoreOutputErr bool
	LinePrefix      string // beginning of the line
	MsgPrefix       string // before the message
	Color           string
	TimeFormat      string
	CallerSkip      int
	dir             string
	dirLen          int
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Caller:     DefaultCaller,
		TimeFormat: DefaultTimeFormat,
		Enable:     true,
		Output: []io.Writer{
			os.Stdout,
		},
	}
}

// NewErrConfig ...
func NewErrConfig() *Config {
	return &Config{
		Caller:     DefaultCaller,
		TimeFormat: DefaultTimeFormat,
		Enable:     true,
		Output: []io.Writer{
			os.Stderr,
		},
	}
}

// Clone ...
func (c *Config) Clone() *Config {
	x := *c
	x.Output = nil
	for _, v := range c.Output {
		x.Output = append(x.Output, v)
	}
	return &x
}

func (c *Config) bufferPrepare() (buf *bytes.Buffer) {

	buf = &bytes.Buffer{}

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
		_, file, line, ok := runtime.Caller(4 + c.CallerSkip)
		if ok {
			file = strings.TrimSuffix(file, hideExt)
			if c.Caller == CallerShort {
				file = filepath.Base(file)
			} else {
				if c.dirLen > 0 && strings.HasPrefix(file, c.dir) {
					file = file[c.dirLen:]
				}
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

func (c *Config) bufferEnd(buf *bytes.Buffer) (n int, err error) {

	if c.Color != `` {
		buf.WriteString("\x1b[0m")
	}

	buf.WriteRune('\n')

	n, err = c.DirectWrite(buf.Bytes())
	buf.Reset()

	return
}

// WriteString ...
func (c *Config) WriteString(msg string) (n int, err error) {

	if c == nil {
		err = ErrNotInited
		return
	}
	if !c.Enable {
		err = ErrDisabled
		return
	}
	if len(c.Output) == 0 {
		err = ErrNoOutput
		return
	}

	buf := c.bufferPrepare()

	size := len(msg)
	empty := true
	if size > 0 {
		for i := size - 1; i >= 0; i-- {
			if msg[i] != '\n' {
				empty = false
				size = i + 1
				break
			}
		}
	}
	if !empty {
		buf.WriteString(msg[:size])
	}

	return c.bufferEnd(buf)
}

// Write ...
func (c *Config) Write(msg []byte) (n int, err error) {

	if c == nil {
		err = ErrNotInited
		return
	}
	if !c.Enable {
		err = ErrDisabled
		return
	}
	if len(c.Output) == 0 {
		err = ErrNoOutput
		return
	}

	buf := c.bufferPrepare()

	size := len(msg)
	empty := true
	if size > 0 {
		for i := size - 1; i >= 0; i-- {
			if msg[i] != '\n' {
				empty = false
				size = i + 1
				break
			}
		}
	}
	if !empty {
		buf.Write(msg[:size])
	}

	return c.bufferEnd(buf)
}

// DirectWrite write raw byte slice to Output, no color/time/file prefix ect.
func (c *Config) DirectWrite(p []byte) (n int, err error) {
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

// StdLogger make a log.Logger
func (c *Config) StdLogger() *log.Logger {
	return log.New(c, ``, 0)
}
