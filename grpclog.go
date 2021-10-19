package zog

import "fmt"

// for grpclog LoggerV2 interface
// https://pkg.go.dev/google.golang.org/grpc/grpclog#LoggerV2

// Info ...
func (i *Logger) Info(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CInfo, s)
}

// Infoln ...
func (i *Logger) Infoln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CInfo, s)
}

// Infof ...
func (i *Logger) Infof(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CInfo, s)
}

// Warning ...
func (i *Logger) Warning(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CWarn, s)
}

// Warningln ...
func (i *Logger) Warningln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CWarn, s)
}

// Warningf ...
func (i *Logger) Warningf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CWarn, s)
}

// Error ...
func (i *Logger) Error(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CError, s)
}

// Errorln ...
func (i *Logger) Errorln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CError, s)
}

// Errorf ...
func (i *Logger) Errorf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CError, s)
}

// Fatal ...
func (i *Logger) Fatal(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CFatal, s)
}

// Fatalln ...
func (i *Logger) Fatalln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CFatal, s)
}

// Fatalf ...
func (i *Logger) Fatalf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CFatal, s)
}

// V ...
func (i *Logger) V(l int) bool {
	return true
}
