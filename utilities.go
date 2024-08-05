package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var _defaultDebugColor color.Color = color.RGBA{G: 100, A: 100}

type debugCache struct {
	cachedImage *ebiten.Image
	debugColor  color.Color
}

func (f *debugCache) Image(w, h int, clr color.Color) *ebiten.Image {
	if f.debugColor == nil {
		f.debugColor = _defaultDebugColor
	}

	if f.cachedImage == nil || !isColorEqual(f.debugColor, clr) {
		f.cachedImage = debugImage(int(w), int(h), clr)
	}

	return f.cachedImage
}

func (f debugCache) Copy() debugCache {
	f.cachedImage = nil
	return f
}

func (f *debugCache) Clean() {
	f.cachedImage = nil
}

func debugImage(w, h int, clr ...color.Color) *ebiten.Image {
	debugColor := _defaultDebugColor
	if len(clr) != 0 && clr[0] != nil {
		debugColor = clr[0]
	}

	if w <= 0 {
		w = 1
	}

	if h <= 0 {
		h = 1
	}

	debugImg := ebiten.NewImage(w, h)
	debugImg.Fill(debugColor)
	return debugImg
}

func isColorEqual(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	return ar == br && ag == bg && ab == bb && aa == ba
}

// var (
// 	_centerImage     = ebiten.NewImage(5, 5)
// 	_centerBaseImage = ebiten.NewImage(5, 5)
// 	_vertexImage     = ebiten.NewImage(3, 3)
// )

// func drawVertexesAndBarycenter(screen *ebiten.Image, ctr controller, vertexes []Vector) {
// 	sync.OnceFunc(func() {
// 		_centerImage.Fill(color.RGBA{R: 255, A: 255})
// 		_centerBaseImage.Fill(color.White)
// 		_vertexImage.Fill(color.White)
// 	})()

// 	mX, mY := ctr.Moved()
// 	// NewImage(_centerBaseImage, AlignCenter).Move(mX, mY).Draw(screen)
// 	NewImage(_centerImage, AlignCenter).Move(mX, mY).Draw(screen)

// 	for _, v := range vertexes {
// 		NewImage(_vertexImage, AlignCenter).Move(v.X, v.Y).Draw(screen)
// 	}
// }
