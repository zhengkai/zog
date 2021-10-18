package zog

import "fmt"

// for grpclog LoggerV2 interface
// https://pkg.go.dev/google.golang.org/grpc/grpclog#LoggerV2

// Info ...
func (i *Input) Info(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CInfo, s)
}

// Infoln ...
func (i *Input) Infoln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CInfo, s)
}

// Infof ...
func (i *Input) Infof(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CInfo, s)
}

// Warning ...
func (i *Input) Warning(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CWarn, s)
}

// Warningln ...
func (i *Input) Warningln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CWarn, s)
}

// Warningf ...
func (i *Input) Warningf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CWarn, s)
}

// Error ...
func (i *Input) Error(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CError, s)
}

// Errorln ...
func (i *Input) Errorln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CError, s)
}

// Errorf ...
func (i *Input) Errorf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CError, s)
}

// Fatal ...
func (i *Input) Fatal(msg ...interface{}) {
	s := fmt.Sprint(msg...)
	i.write(i.CFatal, s)
}

// Fatalln ...
func (i *Input) Fatalln(msg ...interface{}) {
	s := fmt.Sprintln(msg...)
	i.write(i.CFatal, s)
}

// Fatalf ...
func (i *Input) Fatalf(format string, msg ...interface{}) {
	s := fmt.Sprintf(format, msg...)
	i.write(i.CFatal, s)
}

// V ...
func (i *Input) V(l int) bool {
	return true
}
