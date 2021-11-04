package zog

import "fmt"

// for grpclog LoggerV2 interface
// https://pkg.go.dev/google.golang.org/grpc/grpclog#LoggerV2

// Info ...
func (i *Logger) Info(msg ...interface{}) {
	fmt.Fprint(i.GetConfig(i.CInfo), msg...)
}

// Infoln ...
func (i *Logger) Infoln(msg ...interface{}) {
	fmt.Fprintln(i.GetConfig(i.CInfo), msg...)
}

// Infof ...
func (i *Logger) Infof(format string, msg ...interface{}) {
	fmt.Fprintf(i.GetConfig(i.CInfo), format, msg...)
}

// Warning ...
func (i *Logger) Warning(msg ...interface{}) {
	fmt.Fprint(i.GetConfig(i.CWarn), msg...)
}

// Warningln ...
func (i *Logger) Warningln(msg ...interface{}) {
	fmt.Fprintln(i.GetConfig(i.CWarn), msg...)
}

// Warningf ...
func (i *Logger) Warningf(format string, msg ...interface{}) {
	fmt.Fprintf(i.GetConfig(i.CWarn), format, msg...)
}

// Error ...
func (i *Logger) Error(msg ...interface{}) {
	fmt.Fprint(i.GetConfig(i.CError), msg...)
}

// Errorln ...
func (i *Logger) Errorln(msg ...interface{}) {
	fmt.Fprintln(i.GetConfig(i.CError), msg...)
}

// Errorf ...
func (i *Logger) Errorf(format string, msg ...interface{}) {
	fmt.Fprintf(i.GetConfig(i.CError), format, msg...)
}

// Fatal ...
func (i *Logger) Fatal(msg ...interface{}) {
	fmt.Fprint(i.GetConfig(i.CFatal), msg...)
}

// Fatalln ...
func (i *Logger) Fatalln(msg ...interface{}) {
	fmt.Fprintln(i.GetConfig(i.CFatal), msg...)
}

// Fatalf ...
func (i *Logger) Fatalf(format string, msg ...interface{}) {
	fmt.Fprintf(i.GetConfig(i.CFatal), format, msg...)
}

// V ...
func (i *Logger) V(l int) bool {
	return true
}
