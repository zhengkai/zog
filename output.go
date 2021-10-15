package zog

import (
	"io"
	"os"
	"time"
)

type outputTime interface {
	TimeWrite([]byte, time.Time) (int, error)
}

type output struct {
	dst       []io.Writer
	IgnoreErr bool
}

func (o *output) Write(p []byte) (n int, err error) {
	for _, d := range o.dst {
		n, err = d.Write(p)
		if err != nil && !o.IgnoreErr {
			return
		}
	}
	return
}

func (o *output) TimeWrite(p []byte, t time.Time) (n int, err error) {
	for _, d := range o.dst {
		x, ok := d.(outputTime)
		if ok {
			n, err = x.TimeWrite(p, t)
		} else {
			n, err = d.Write(p)
		}
		if err != nil && !o.IgnoreErr {
			return
		}
	}
	return
}

// OutputFile ...
func OutputFile(name string, isAppend bool) (f *os.File, err error) {

	flag := os.O_WRONLY | os.O_CREATE
	if isAppend {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}

	f, err = os.OpenFile(name, flag, 0644)
	if err != nil {
		return
	}

	return
}
