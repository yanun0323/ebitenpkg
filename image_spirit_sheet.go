package ebitenpkg

import (
	sysimage "image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Direction Align

type SpiritSheetOption func(direction Direction, imgWidth, imgHeight int, currentSecond int) sysimage.Rectangle

func (opt *SpiritSheetOption) subImage(img *ebiten.Image, direction Direction, currentSecond int) *ebiten.Image {
	if opt == nil {
		return img
	}

	fn := *opt

	rect := fn(direction, img.Bounds().Dx(), img.Bounds().Dy(), currentSecond)

	return img.SubImage(rect).(*ebiten.Image)
}
