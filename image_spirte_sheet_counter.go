package ebitenpkg

import (
	"image"
)

type spriteSheetOptionCounter struct {
	columnCount int
	rowCount    int
	maxIndex    int
	execute     func(fps, timestamp int, direction Direction) (index int, scaleX, scaleY int)
}

func SpriteSheetOptionCounter(
	columnCount, rowCount, maxIndex int,
	handler func(fps, timestamp int, direction Direction) (index int, scaleX, scaleY int),
) SpriteSheetOption {
	if maxIndex <= 0 {
		maxIndex = columnCount * rowCount
	}

	return &spriteSheetOptionCounter{
		columnCount: max(columnCount, 1),
		rowCount:    max(rowCount, 1),
		maxIndex:    max(maxIndex, 1),
		execute:     handler,
	}
}

func (opt *spriteSheetOptionCounter) Mask(src image.Image, fps, timestamp int, direction Direction) (image.Rectangle, int, int) {
	if opt.execute == nil {
		return src.Bounds(), 1, 1
	}

	bounds := src.Bounds()
	sizeW, sizeH := bounds.Dx()/opt.columnCount, bounds.Dy()/opt.rowCount
	idx, sX, sY := opt.execute(fps, timestamp, direction)
	idx = idx % opt.maxIndex

	start := image.Point{
		X: sizeW * (idx % opt.columnCount),
		Y: sizeH * (idx / opt.columnCount),
	}

	end := image.Point{
		X: start.X + sizeW,
		Y: start.Y + sizeH,
	}

	return image.Rectangle{Min: start, Max: end}, sX, sY
}
