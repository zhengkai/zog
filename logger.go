package zog

import (
	"log"
)

// Logger ...
type Logger struct {
	CDefault *Config
	CPrint   *Config
	CDebug   *Config
	CInfo    *Config
	CError   *Config
	CWarn    *Config
	CFatal   *Config
	CWatch   *Config
}

// NewSimple ...
func NewSimple(file string) (i *Logger, err error) {

	f, err := NewFile(file, false)
	if err != nil {
		return
	}

	cfg := NewConfig()
	cfg.Output = append(cfg.Output, f)

	i = &Logger{
		CDefault: cfg,
	}

	return
}

// Writer for io.Writer
func (i *Logger) Write(p []byte) (n int, err error) {
	return i.CDefault.Write(p)
}

func (i *Logger) write(c *Config, p string) {
	i.GetConfig(c).WriteString(p)
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
	return []*Config{i.CDefault, i.CPrint, i.CDebug, i.CInfo, i.CError, i.CWarn, i.CFatal}
}

// SetDirPrefix dir prefix in caller filename will be hidden
func (i *Logger) SetDirPrefix(d string) {
	for _, c := range i.AllConfig() {
		if c != nil {
			c.SetDirPrefix(d)
			c.Caller = CallerLong
		}
	}
}

// StdLogger make a log.Logger
func (i *Logger) StdLogger() *log.Logger {
	return log.New(i, ``, 0)
}
