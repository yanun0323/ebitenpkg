package ebito

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image struct {
	img *ebiten.Image
	*drawOption
}

func NewImageFromImage(img image.Image, a ...Align) *Image {
	return NewImage(ebiten.NewImageFromImage(img), a...)
}

func NewImage(img *ebiten.Image, a ...Align) *Image {
	return &Image{
		img:        img,
		drawOption: NewDrawOption(float64(img.Bounds().Dx()), float64(img.Bounds().Dy()), a...),
	}
}

func (f Image) Image() *ebiten.Image {
	return f.img
}

func (f Image) Copy() *Image {
	b := f.img.Bounds()
	f.img = ebiten.NewImage(b.Dx(), b.Dy())
	f.img.DrawImage(f.img, nil)
	f.drawOption = f.drawOption.copy()
	return &f
}

// Draw is an alias to screen.DrawImage(img.Image(), img.Option())
func (f Image) Draw(screen *ebiten.Image) {
	screen.DrawImage(f.img, f.Option())
}

func (f Image) DebugDraw(screen *ebiten.Image, borderWidth ...int) {
	f.Draw(screen)
	screen.DrawImage(DebugImageFromImage(f.img, borderWidth...), f.debugOption(borderWidth...))
}

func (f *Image) UpdateImageFromImage(img image.Image) {
	f.UpdateImage(ebiten.NewImageFromImage(img))
}

func (f *Image) UpdateImage(img *ebiten.Image) {
	f.img = img
	f.recalculate()
}

func (f *Image) recalculate() {
	bounds := f.img.Bounds()
	f.updateReference(float64(bounds.Dx()), float64(bounds.Dy()))
}

// Extension of DrawOption

func (f Image) WithMovement(x, y float64) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withMovement(x, y)
	return ff
}

func (f Image) WithScaleRatio(x, y float64) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withScaleRatio(x, y)
	return ff
}

func (f Image) WithRotation(degree float64) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withRotation(degree)
	return ff
}

func (f Image) WithAlignment(a Align) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withAlignment(a)
	return ff
}
