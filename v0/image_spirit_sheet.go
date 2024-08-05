package ebitenpkg

import (
	sysimage "image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheetOption struct {
	SpriteWidth  int
	SpriteHeight int
	Fn           func(direction Direction, spriteWidth, spriteHeight int, currentSecond int64) sysimage.Rectangle
}

func (opt SpriteSheetOption) subImage(img *ebiten.Image, direction Direction, currentSecond int64) *ebiten.Image {
	if opt.Fn == nil {
		return img
	}

	rect := opt.Fn(direction, img.Bounds().Dx(), img.Bounds().Dy(), currentSecond)

	return img.SubImage(rect).(*ebiten.Image)
}
