package stringctr

import (
	"fmt"
	"testing"
)

func Test_Prettier_IgnoreEmptyIsTrue(t *testing.T) {
	expected := `***   Hello, Golang worl   ***
***   d and goodbye, Cla   ***
***   ng world !!          ***
`
	prettier := NewPrettier()
	actual, err := prettier.
		Target("Hello, Golang world and goodbye, Clang world !!").
		RowLimit(30).
		TargetLimit(18).
		Margin(3).
		Padding("*").
		IgnoreEmpty(true).
		Prettify()

	fmt.Printf("%s", actual)
	if err != nil {
		t.Error("something went wrong")
	}
	if actual != expected {
		t.Error("actual is unexpected")
	}
}

/*
func Test_Prettier_IgnoreEmptyIsFalse(t *testing.T) {
	expected := `***   Hello, Golang        ***
***   world and goodbye,   ***
***   Clang world !!       ***
`
	prettier := NewPrettier()
	actual, err := prettier.
		Target("Hello, Golang world and goodbye, Clang world !!").
		RowLimit(30).
		TargetLimit(18).
		Margin(3).
		Padding("*").
		IgnoreEmpty(false).
		Prettify()

	fmt.Printf("%s", actual)
	if err != nil {
		t.Error("something went wrong")
	}
	if actual != expected {
		t.Error("actual is unexpected")
	}
}
*/
