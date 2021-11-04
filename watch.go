package zog

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
)

var errWatch = errors.New(`Incorrect use of props`)

// Watch ...
func (i *Logger) Watch(err *error, prefix ...interface{}) {
	i.watch(err, false, prefix...)
}

// WatchStack ...
func (i *Logger) WatchStack(err *error, prefix ...interface{}) {
	i.watch(err, true, prefix...)
}

func (i *Logger) watch(err *error, stack bool, prefix ...interface{}) {

	if err == nil {
		err = &errWatch
	} else if *err == nil {
		return
	}

	var buf bytes.Buffer

	if len(prefix) > 0 {
		s := fmt.Sprintln(prefix...)
		buf.WriteString(s[:len(s)-1])
		buf.WriteRune(' ')
	}

	buf.WriteString(getFrame(2).Function)
	buf.WriteString(`() `)
	buf.WriteString((*err).Error())

	cfg := i.GetConfig(i.CError)

	if stack {
		a := bytes.Split(debug.Stack(), []byte{'\n'})
		if len(a) > 8 {
			// pretty stack
			var prefix []byte
			if cfg.dirLen > 0 {
				prefix = []byte("\t" + cfg.dir)
			}
			for _, v := range a[7 : len(a)-1] {
				buf.Write([]byte{'\n', '\t'})
				if prefix != nil && bytes.HasPrefix(v, prefix) {
					v = v[len(prefix)-1:]
					v[0] = '\t'
				}
				buf.Write(v)
			}
		}
	}

	buf.WriteTo(cfg)
}

func getFrame(skipFrames int) (f runtime.Frame) {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	maxIdx := skipFrames + 2

	// Set size to maxIdx+2 to ensure we have room for one more caller than we need
	pc := make([]uintptr, maxIdx+2)
	n := runtime.Callers(0, pc)

	f = runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(pc[:n])
		for idx := 0; idx <= maxIdx; idx++ {
			nf, more := frames.Next()
			f = nf
			if !more {
				break
			}
		}
	}

	return
}
