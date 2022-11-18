package main

import (
	"fmt"
	"github.com/ughvj/gotils/src/stringctr"
)

func main() {
	base := "HELLO"
	digit := len(base) + 10;
	pr := stringctr.PaddingRight(base, digit, "O")
	pl := stringctr.PaddingLeft(base, digit, "H")
	pb := stringctr.PaddingBoth(base, " HELLO ")
	pbwls := stringctr.PaddingBothWithLineSymmetry(base, " HELLO ")

	fmt.Println("PaddingRight: " + pr)
	fmt.Println("PaddingLeft: " + pl)
	fmt.Println("PaddingBoth: " + pb)
	fmt.Println("PaddingBothWithLineSymmetry: " + pbwls);
}
