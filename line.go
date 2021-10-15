package zog

import "fmt"

func defaultLine(time, prefix, caller, msg string) string {
	return fmt.Sprintf(`%s%s %s %s`, time, prefix, caller, msg)
}
