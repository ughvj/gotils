package stringctr

import (
	"testing"

	"github.com/ughvj/gotils/src/testmsg"
)

func Test_PaddingLeftWillSucceed(t *testing.T) {
	// Given.
	target := "hogehoge"

	expected := []string{
		"hogehoge",
		"hogehoge",
		"hogehoge",
		"hogehoge",
		"hogehoge",
		"0hogehoge",
		"xyzhogehoge",
		"xyzxhogehoge",
		// "あいうえおあhogehoge", // unsupported.
	}

	// When.
	actual := []string{
		PaddingLeft(target, -1, "0"),
		PaddingLeft(target, 0, "0"),
		PaddingLeft(target, 7, "0"),
		PaddingLeft(target, 8, "0"),
		PaddingLeft(target, 8, ""),
		PaddingLeft(target, 9, "0"),
		PaddingLeft(target, 11, "xyz"),
		PaddingLeft(target, 12, "xyz"),
		// PaddingLeft(target, 14, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}

func Test_PaddingRightWillSucceed(t *testing.T) {
	// Given.
	target := "hogehoge"

	expected := []string{
		"hogehoge",
		"hogehoge",
		"hogehoge",
		"hogehoge",
		"hogehoge",
		"hogehoge0",
		"hogehogexyz",
		"hogehogexyzx",
		// "hogehogeあいうえおあ", // unsupported.
	}

	// When.
	actual := []string{
		PaddingRight(target, -1, "0"),
		PaddingRight(target, 0, "0"),
		PaddingRight(target, 7, "0"),
		PaddingRight(target, 8, "0"),
		PaddingRight(target, 8, ""),
		PaddingRight(target, 9, "0"),
		PaddingRight(target, 11, "xyz"),
		PaddingRight(target, 12, "xyz"),
		// PaddingRight(target, 14, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}

func Test_PaddingBothWillSucceed(t *testing.T) {
	// Given.
	target := "hogehoge"

	expected := []string{
		"hogehoge",
		"xhogehogex",
		"xxhogehogexx",
		"xyzhogehogexyz",
		// "あいうえおhogehogeあいうえお", // unsupported.
	}

	// When.
	actual := []string{
		PaddingBoth(target, ""),
		PaddingBoth(target, "x"),
		PaddingBoth(target, "xx"),
		PaddingBoth(target, "xyz"),
		// PaddingBoth(target, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}

func Test_PaddingBothWithLineSymmetryWillSucceed(t *testing.T) {
	// Given.
	target := "hogehoge"

	expected := []string{
		"hogehoge",
		"xhogehogex",
		"zyxhogehogexyz",
		// "おえういあhogehogeあいうえお", // unsupported.
	}

	// When.
	actual := []string{
		PaddingBothWithLineSymmetry(target, ""),
		PaddingBothWithLineSymmetry(target, "x"),
		PaddingBothWithLineSymmetry(target, "xyz"),
		// PaddingBothWithLineSymmetry(target, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}
