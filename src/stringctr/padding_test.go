package stringctr

import (
	"testing"

	"github.com/ughvj/gotils/src/testmsg"
)

func Test_LeftWillSucceed(t *testing.T) {
	// Given.
	padding := Padding{}
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
		padding.Left(target, -1, "0", false),
		padding.Left(target, 0, "0", false),
		padding.Left(target, 7, "0", false),
		padding.Left(target, 8, "0", false),
		padding.Left(target, 8, "", false),
		padding.Left(target, 9, "0", false),
		padding.Left(target, 11, "xyz", false),
		padding.Left(target, 12, "xyz", false),
		// padding.Left(target, 14, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}

func Test_RightWillSucceed(t *testing.T) {
	// Given.
	padding := Padding{}
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
		padding.Right(target, -1, "0", false),
		padding.Right(target, 0, "0", false),
		padding.Right(target, 7, "0", false),
		padding.Right(target, 8, "0", false),
		padding.Right(target, 8, "", false),
		padding.Right(target, 9, "0", false),
		padding.Right(target, 11, "xyz", false),
		padding.Right(target, 12, "xyz", false),
		// padding.Right(target, 14, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}

func Test_BothWillSucceed(t *testing.T) {
	// Given.
	padding := Padding{}
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
		padding.Both(target, ""),
		padding.Both(target, "x"),
		padding.Both(target, "xx"),
		padding.Both(target, "xyz"),
		// padding.Both(target, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}

func Test_BothWithLineSymmetryWillSucceed(t *testing.T) {
	// Given.
	padding := Padding{}
	target := "hogehoge"

	expected := []string{
		"hogehoge",
		"xhogehogex",
		"zyxhogehogexyz",
		// "おえういあhogehogeあいうえお", // unsupported.
	}

	// When.
	actual := []string{
		padding.BothWithLineSymmetry(target, ""),
		padding.BothWithLineSymmetry(target, "x"),
		padding.BothWithLineSymmetry(target, "xyz"),
		// padding.BothWithLineSymmetry(target, "あいうえお") // unsupported.
	}

	// Then.
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Error(testmsg.Inspect(expected[i], actual[i]))
		}
	}
}
