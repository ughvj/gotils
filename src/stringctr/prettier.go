/*
 * 文字列を指定した文字数ごとに改行し、装飾するプログラム
 */

package stringctr

import (
	"errors"
)

 /*
  * 並び替えるもとの文字列と、
  * 付け足したり、並び替えるための値を格納する構造体
  */
type Prettier struct {
	target string // 並び替える文字列
	targetLimit int // 文字列の最大値
	rowLimit int // 文字列の最小値
	margin int // 左右の空白
	ignoreEmpty bool // 左右に空白を開けるか
	padding string // 左右に付け足す文字列
	padder PaddingInterface // 文字列を付け足すためのインターフェース
}

/*
 * 新しい Prettier struct を作成する
 */
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

/*
 * target を値を設定
 */
func (p *Prettier) Target(target string) *Prettier {
	p.target = target
	return p
}

/*
 * targetLimit の値を設定
 */
func (p *Prettier) TargetLimit(targetLimit int) *Prettier {
	p.targetLimit = targetLimit
	return p
}

/*
 * rowLimit の値を設定
 */
func (p *Prettier) RowLimit(rowLimit int) *Prettier {
	p.rowLimit = rowLimit
	return p
}

/*
 * margin の値を設定
 */
func (p *Prettier) Margin(margin int) *Prettier {
	p.margin = margin
	return p
}

/*
 * ignoreEmpty の値を設定
 */
func (p *Prettier) IgnoreEmpty(ignoreEmpty bool) *Prettier {
	p.ignoreEmpty = ignoreEmpty
	return p
}

/*
 * padding の値を設定
 */
func (p *Prettier) Padding(padding string) *Prettier {
	p.padding = padding
	return p
}

/*
 * Prettier sturct に設定した値をもとに文字列を整える
 */
func (p *Prettier) Prettify() (prettified string, err error) {
	// target の文字数が 1 未満だった場合エラー
	if len(p.target) < 1 {
		return "", errors.New("target string length must be greater than 0")
	}

	// rowLimit が 1 未満だった場合エラー
	if p.rowLimit < 1 {
		return p.target, errors.New("row limit must be greater than 0")
	}

	// rowLimit が 80 を超える場合エラー
	if p.rowLimit > 80 {
		return p.target, errors.New("row limit must be lower than 80")
	}

	// targetLimit が 1 未満だった場合エラー
	if p.targetLimit < 1 {
		return p.target, errors.New("target limit must be greater than 0")
	}

	// targetLimit が 79 を超える場合エラー
	if p.targetLimit > 79 {
		return p.target, errors.New("target limit must be lower than 79")
	}

	// targetLimit が rowLimit を超える場合エラー
	if p.targetLimit > p.rowLimit {
		return p.target, errors.New("target limit must be lower than row limit")
	}

	// (paddingの長さ * 2 + margin * 2) の値が (rowLimit - targetLimit) の値を超える場合エラー
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
