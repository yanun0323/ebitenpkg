package ebitenpkg

import "image"

type spriteableOptionFunc struct {
	fn func(src image.Image, fps, timestamp int, direction Direction) (mask image.Rectangle, scaleX, scaleY int)
}

func SpriteableOptionFunc(fn func(src image.Image, fps, timestamp int, direction Direction) (mask image.Rectangle, scaleX, scaleY int)) SpriteableOption {
	return &spriteableOptionFunc{fn: fn}
}

func (opt *spriteableOptionFunc) Mask(src image.Image, fps, timestamp int, direction Direction) (image.Rectangle, int, int) {
	return opt.fn(src, fps, timestamp, direction)
}
