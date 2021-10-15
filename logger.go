package zog

import "log"

// MakeLogger make a log.Logger
func (i *Input) MakeLogger() *log.Logger {
	return log.New(i, ``, 0)
}
