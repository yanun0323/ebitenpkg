package ebitenpkg

import "image"

type spriteableOptionRect struct {
	rect image.Rectangle
	fn   func(fps, timestamp int, direction Direction) (offsetX, offsetY, scaleX, scaleY int)
}

func SpriteableOptionRect(rect image.Rectangle, fn func(fps, timestamp int, direction Direction) (offsetX, offsetY, scaleX, scaleY int)) SpriteableOption {
	return &spriteableOptionRect{
		rect: rect,
		fn:   fn,
	}
}

func (opt *spriteableOptionRect) Mask(src image.Image, fps, timestamp int, direction Direction) (image.Rectangle, int, int) {
	oX, oY, sX, sY := opt.fn(fps, timestamp, direction)
	start := image.Point{
		X: opt.rect.Dx() * oX,
		Y: opt.rect.Dy() * oY,
	}

	return image.Rectangle{
		Min: start,
		Max: image.Point{
			X: start.X + opt.rect.Dx(),
			Y: start.Y + opt.rect.Dy(),
		},
	}, sX, sY
}
