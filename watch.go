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
func (i *Input) Watch(err *error, prefix ...interface{}) {
	i.watch(err, false, prefix...)
}

// WatchStack ...
func (i *Input) WatchStack(err *error, prefix ...interface{}) {
	i.watch(err, true, prefix...)
}

func (i *Input) watch(err *error, stack bool, prefix ...interface{}) {

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

	cfg := i.CError
	if cfg == nil {
		cfg = i.CDefault
	}

	if stack {
		a := bytes.Split(debug.Stack(), []byte{'\n'})
		if len(a) > 7 {
			// pretty stack
			var prefix []byte
			if cfg.dirLen > 0 {
				prefix = []byte("\t" + cfg.dir)
			}
			fmt.Println(string(prefix))
			for _, v := range a[7:] {
				buf.Write([]byte{'\n', '\t'})
				if prefix != nil && bytes.HasPrefix(v, prefix) {
					v = v[len(prefix)-1:]
					v[0] = '\t'
				}
				buf.Write(v)
			}
		}
	}

	i.write(i.CError, buf.String())
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
		for more, idx := true, 0; more && idx <= maxIdx; idx++ {
			var nf runtime.Frame
			nf, more = frames.Next()
			if idx == maxIdx {
				f = nf
			}
		}
	}

	return f
}
