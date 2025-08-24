package ebitenpkg

import "image"

type spriteSheetOptionFunc struct {
	fn func(src image.Image, fps, timestamp int, direction Direction) (mask image.Rectangle, scaleX, scaleY int)
}

func SpriteSheetOptionFunc(fn func(src image.Image, fps, timestamp int, direction Direction) (mask image.Rectangle, scaleX, scaleY int)) SpritSheetOption {
	return &spriteSheetOptionFunc{fn: fn}
}

func (opt *spriteSheetOptionFunc) Mask(src image.Image, fps, timestamp int, direction Direction) (image.Rectangle, int, int) {
	return opt.fn(src, fps, timestamp, direction)
}
