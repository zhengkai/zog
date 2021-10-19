package zog

// just like Lshortfile/Llongfile in pkg/log
const (
	CallerNone  = CallerType(iota)
	CallerShort // caller:42
	CallerLong  // /dir/caller:42
)

// CallerType ...
type CallerType uint8

const hideExt = `.go`
