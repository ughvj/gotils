package testmsg

import "fmt"

// Inspect is inspect variables in tests when will be error.
func Inspect(expected interface{}, actual interface{}) string {
	return fmt.Sprintf("\nExpected: %#v\nActual  : %#v", expected, actual)
}
