package zog

// just like Lshortfile/Llongfile in pkg/log
const (
	CallerNone    = callerType(iota + 1)
	CallerShort   // caller.go:42
	CallerShorter // caller:42
	CallerLong    // /dir/caller.go:42
)

type callerType uint8
