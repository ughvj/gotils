package stringctr

import (
	"errors"
)

type Prettier struct {
	target string
	targetLimit int
	rowLimit int
	margin int
	ignoreEmpty bool
	padding string
	padder PaddingInterface
}

func NewPrettier() *Prettier {
	return &Prettier{
		target: "",
		targetLimit: 20,
		rowLimit: 40,
		margin: 1,
		ignoreEmpty: true,
		padding: "*",
		padder: NewPadding(),
	}
}

func (p *Prettier) Target(target string) *Prettier {
	p.target = target
	return p
}

func (p *Prettier) TargetLimit(targetLimit int) *Prettier {
	p.targetLimit = targetLimit
	return p
}

func (p *Prettier) RowLimit(rowLimit int) *Prettier {
	p.rowLimit = rowLimit
	return p
}

func (p *Prettier) Margin(margin int) *Prettier {
	p.margin = margin
	return p
}

func (p *Prettier) IgnoreEmpty(ignoreEmpty bool) *Prettier {
	p.ignoreEmpty = ignoreEmpty
	return p
}

func (p *Prettier) Padding(padding string) *Prettier {
	p.padding = padding
	return p
}

func (p *Prettier) Prettify() (prettified string, err error) {
	if len(p.target) < 1 {
		return "", errors.New("target string length must be greater than 0")
	}

	if p.rowLimit < 1 {
		return p.target, errors.New("row limit must be greater than 0")
	}
	if p.rowLimit > 80 {
		return p.target, errors.New("row limit must be lower than 80")
	}
	if p.targetLimit < 1 {
		return p.target, errors.New("target limit must be greater than 0")
	}
	if p.targetLimit > 79 {
		return p.target, errors.New("target limit must be lower than 79")
	}
	if p.targetLimit > p.rowLimit {
		return p.target, errors.New("target limit must be lower than row limit")
	}

	if len(p.padding) * 2 + p.margin * 2 > p.rowLimit - p.targetLimit {
		return p.target, errors.New("style settings is invalid")
	}

	for i := 1; i <= len(p.target); i++ {
		segment := ""

		if i % p.targetLimit == 0 {
			segment = p.target[i - p.targetLimit:i]
		} else if i == len(p.target) {
			segment = p.padder.Right(p.target[i - (i % p.targetLimit):i], p.targetLimit, " ", false)
		} else {
			continue
		}
	
		for j := 0; j < p.margin; j++ {
			segment = p.padder.Both(segment, " ")
		}
		segment = p.padder.Left(segment, len(segment) + (p.rowLimit - len(segment)) / 2, p.padding, false)
		segment = p.padder.Right(segment, p.rowLimit, p.padding, true)

		prettified += segment + "\n"
	}

	return prettified, nil
}
