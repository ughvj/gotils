package stringctr

import (
	"testing"

	"github.com/ughvj/gotils/src/testmsg"
)

func Test_PadByWillSucceed(t *testing.T) {
	// Given.
	target := "hogehoge"

	e1 := "hogehoge"
	// e2 := "hogehoge"
	// e3 := "hogehoge"
	// e4 := "hogehoge"
	// e5 := "0hogehoge"
	// e6 := "xyzhogehoge"
	// e7 := "xyzxhogehoge"

	// When.
	a1 := PadBy(target, -1, "0")
	// a2 := PadBy(target, 0, "0")
	// a3 := PadBy(target, 7, "0")
	// a4 := PadBy(target, 8, "0")
	// a5 := PadBy(target, 9, "0")
	// a6 := PadBy(target, 11, "xyz")
	// a7 := PadBy(target, 12, "xyz")

	// Then.
	if e1 != a1 {
		t.Error(testmsg.Inspect(e1, a1))
	}
}
