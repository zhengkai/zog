package zog

import "os"

// Default value for NewConfig / NewErrConfig
var (
	DefaultCaller     = CallerShort
	DefaultTimeFormat = TimeYear
)

// Default permission when create file / dir
var (
	DefaultPermFile os.FileMode = 0644
	DefaultPermDir  os.FileMode = 0755
)
