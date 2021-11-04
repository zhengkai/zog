package zog

import (
	"io"
	"log"
)

// Logger ...
type Logger struct {
	CDefault *Config
	CWrite   *Config
	CPrint   *Config
	CDebug   *Config
	CInfo    *Config
	CError   *Config
	CWarn    *Config
	CFatal   *Config
}

// NewSimple ...
func NewSimple(file string) (i *Logger, err error) {

	f, err := NewFile(file, false)
	if err != nil {
		return
	}

	cfg := NewConfig()
	cfg.AddOutput(f)

	i = &Logger{
		CDefault: cfg,
	}

	return
}

// Writer for io.Writer
func (i *Logger) Write(p []byte) (n int, err error) {
	c := i.CWrite
	if c == nil {
		c = i.CDefault
	}
	n, err = c.Write(p)
	return
}

func (i *Logger) write(c *Config, p string) {
	if c == nil {
		c = i.CDefault
	}
	io.WriteString(c, p)
}

// GetConfig ...
func (i *Logger) GetConfig(c *Config) *Config {
	if c != nil {
		return c
	}
	return i.CDefault
}

// AllConfig ...
func (i *Logger) AllConfig() []*Config {
	return []*Config{i.CDefault, i.CWrite, i.CPrint, i.CDebug, i.CInfo, i.CError, i.CWarn, i.CFatal}
}

// SetDirPrefix dir prefix in caller filename with be hidden
func (i *Logger) SetDirPrefix(d string) {
	for _, c := range i.AllConfig() {
		if c != nil {
			c.SetDirPrefix(d)
		}
	}
}

// StdLogger make a log.Logger
func (i *Logger) StdLogger() *log.Logger {
	return log.New(i, ``, 0)
}
