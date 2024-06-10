package ebitenpkg

import "github.com/hajimehoshi/ebiten/v2/text/v2"

type Align int

const (
	AlignCenter Align = iota
	AlignTopCenter
	AlignBottomCenter
	AlignLeading
	AlignTopLeading
	AlignBottomLeading
	AlignTrailing
	AlignTopTrailing
	AlignBottomTrailing
)

func (a Align) TextAlign() text.Align {
	switch a {
	case AlignLeading:
		return text.AlignStart
	case AlignCenter:
		return text.AlignCenter
	case AlignTrailing:
		return text.AlignEnd
	default:
		return text.AlignStart
	}
}
