package zog

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
)

var errWatch = errors.New(`Incorrect use of props`)

// Watch ...
func (i *Input) Watch(err *error, prefix ...interface{}) {

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

	buf.WriteString(getFrame(1).Function)
	buf.WriteString(`() `)
	buf.WriteString((*err).Error())

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
