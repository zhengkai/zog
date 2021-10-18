package zog

// just like Lshortfile/Llongfile in pkg/log
const (
	CallerNone  = CallerType(iota)
	CallerShort // caller.go:42
	CallerLong  // /dir/caller.go:42
)

// CallerType ...
type CallerType uint8
