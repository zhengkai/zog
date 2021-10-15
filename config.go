package zog

import "bytes"

type config struct {
	o          *output
	caller     callerType
	timeFormat string
}

func (c *config) write(msg string) {

	var buf bytes.Buffer
	buf.WriteString(msg)

	buf.WriteTo(c.o)
}
