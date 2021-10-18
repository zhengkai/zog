package zog

import (
	"io"
	"os"
)

const baseFileMode = os.O_WRONLY | os.O_CREATE

// Output ...
type Output struct {
	Dest      []io.Writer
	IgnoreErr bool
}

func (o *Output) Write(p []byte) (n int, err error) {
	for _, d := range o.Dest {
		n, err = d.Write(p)
		if err != nil && !o.IgnoreErr {
			return
		}
	}
	return
}

// OutputFile ...
func OutputFile(name string, isAppend bool) (f *os.File, err error) {
	flag := baseFileMode
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
