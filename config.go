package zog

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Config ...
type Config struct {
	Output     *Output
	Caller     callerType
	TimeFormat string
}

func (c *Config) writePrepare() *bytes.Buffer {

	var buf bytes.Buffer

	if c.TimeFormat != `` {
		buf.WriteString(time.Now().Format(c.TimeFormat))
	}

	if c.Caller != CallerNone {
		_, file, line, ok := runtime.Caller(3)
		if ok {
			if c.Caller == CallerShorter {
				file = strings.TrimSuffix(file, `.go`)
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

	return &buf
}

func (c *Config) writeAB(msg []byte) {

	buf := c.writePrepare()

	(*buf).Write(msg)
	l := len(msg)
	if l == 0 || msg[l-1] != '\n' {
		(*buf).WriteRune('\n')
	}

	(*buf).WriteTo(c.Output)
}

func (c *Config) write(msg string) {

	buf := c.writePrepare()

	(*buf).WriteString(msg)
	if !strings.HasSuffix(msg, "\n") {
		(*buf).WriteRune('\n')
	}

	(*buf).WriteTo(c.Output)
}
