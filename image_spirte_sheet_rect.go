package ebitenpkg

import "image"

type spriteSheetOptionRect struct {
	rect image.Rectangle
	fn   func(fps, timestamp int, direction Direction) (offsetX, offsetY, scaleX, scaleY int)
}

func SpriteSheetOptionRect(rect image.Rectangle, fn func(fps, timestamp int, direction Direction) (offsetX, offsetY, scaleX, scaleY int)) SpriteSheetOption {
	return &spriteSheetOptionRect{
		rect: rect,
		fn:   fn,
	}
}

func (opt *spriteSheetOptionRect) Mask(src image.Image, fps, timestamp int, direction Direction) (image.Rectangle, int, int) {
	oX, oY, sX, sY := opt.fn(fps, timestamp, direction)
	start := image.Point{
		X: opt.rect.Min.X + opt.rect.Dx()*oX,
		Y: opt.rect.Min.Y + opt.rect.Dy()*oY,
	}

	return image.Rectangle{
		Min: start,
		Max: image.Point{
			X: start.X + opt.rect.Dx(),
			Y: start.Y + opt.rect.Dy(),
		},
	}, sX, sY
}
