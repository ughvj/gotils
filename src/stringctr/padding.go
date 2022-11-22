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
func (p *Padding) Left(target string, digit int, pad string, reverse bool) string {
	return padBy(target, digit, pad, reverse, "left")
}

// PaddingRight is pad the string on the right by any character except for multi-byte.
func (p *Padding) Right(target string, digit int, pad string, reverse bool) string {
	return padBy(target, digit, pad, reverse, "right")
}

// PaddingBoth is pad the string on both(left side + right side).
func (p *Padding) Both(target string, pad string) string {
	padded := p.Left(target, len(target)+len(pad), pad, false)
	return p.Right(padded, len(padded)+len(pad), pad, false)
}

// PaddingBothWithLineSymmetry is pad the string on both(left side(reversed) + right side).
func (p *Padding) BothWithLineSymmetry(target string, pad string) string {
	padded := p.Left(target, len(target)+len(pad), pad, true)
	return p.Right(padded, len(padded)+len(pad), pad, false)
}

func NewPadding() *Padding {
	return &Padding{}
}

type Padding struct {}

type PaddingInterface interface {
	Left(target string, digit int, pad string, reverse bool) string
	Right(target string, digit int, pad string, reverse bool) string
	Both(target string, pad string) string
	BothWithLineSymmetry(target string, pad string) string
}
