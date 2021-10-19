package zog

import "fmt"

// Print ...
func (i *Logger) Print(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CPrint, s)
}

// Printf ...
func (i *Logger) Printf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CPrint, s)
}

// Println ...
func (i *Logger) Println(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CPrint, s)
}

// Debug ...
func (i *Logger) Debug(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CDebug, s)
}

// Debugf ...
func (i *Logger) Debugf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CDebug, s)
}

// Debugln ...
func (i *Logger) Debugln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CDebug, s)
}
