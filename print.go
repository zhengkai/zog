package zog

import "fmt"

// Print ...
func (i *Logger) Print(msg ...interface{}) {
	fmt.Fprint(i.GetConfig(i.CPrint), msg...)
}

// Printf ...
func (i *Logger) Printf(format string, msg ...interface{}) {
	fmt.Fprintf(i.GetConfig(i.CPrint), format, msg...)
}

// Println ...
func (i *Logger) Println(msg ...interface{}) {
	fmt.Fprintln(i.GetConfig(i.CPrint), msg...)
}

// Debug ...
func (i *Logger) Debug(msg ...interface{}) {
	fmt.Fprint(i.GetConfig(i.CDebug), msg...)
}

// Debugf ...
func (i *Logger) Debugf(format string, msg ...interface{}) {
	fmt.Fprintf(i.GetConfig(i.CDebug), format, msg...)
}

// Debugln ...
func (i *Logger) Debugln(msg ...interface{}) {
	fmt.Fprintln(i.GetConfig(i.CDebug), msg...)
}
