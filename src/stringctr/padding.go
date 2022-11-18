package stringctr

func padBy(target string, digit int, pad string, reversed bool, direction string) (padded string) {
	lack := digit - len(target)

	if lack <= 0 {
		return target
	}

	if len(pad) == 0 {
		return target
	}

	loop := lack / len(pad)
	segment := lack % len(pad)

	pads := ""
	for i := 0; i < loop; i++ {
		pads = pads + pad
	}

	padding := pads + pad[:segment]
	switch direction {
	case "left":
		if reversed {
			padded = reverse(padding) + target
		} else {
			padded = padding + target
		}
	case "right":
		if reversed {
			padded = target + reverse(padding)
		} else {
			padded = target + padding
		}
	default:
		padded = target
	}
	return
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// PaddingLeft is pad the string on the left by any character except for multi-byte.
func PaddingLeft(target string, digit int, pad string) string {
	return padBy(target, digit, pad, false, "left")
}

// PaddingRight is pad the string on the right by any character except for multi-byte.
func PaddingRight(target string, digit int, pad string) string {
	return padBy(target, digit, pad, false, "right")
}

// PaddingBoth is pad the string on both(left side + right side).
func PaddingBoth(target string, pad string) string {
	padded := PaddingLeft(target, len(target)+len(pad), pad)
	return PaddingRight(padded, len(padded)+len(pad), pad)
}

// PaddingBothWithLineSymmetry is pad the string on both(left side(reversed) + right side).
func PaddingBothWithLineSymmetry(target string, pad string) string {
	padded := padBy(target, len(target)+len(pad), pad, true, "left")
	return PaddingRight(padded, len(padded)+len(pad), pad)
}
