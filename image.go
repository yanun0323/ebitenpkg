package ebitenpkg

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
		drawOption: newDrawOption(float64(img.Bounds().Dx()), float64(img.Bounds().Dy()), a...),
	}
}

func (f Image) Image() *ebiten.Image {
	return f.img
}

func (f Image) Copy() *Image {
	b := f.img.Bounds()
	img := ebiten.NewImage(b.Dx(), b.Dy())
	img.DrawImage(f.img, nil)
	f.img = img
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

func (f Image) WithOption(opt DrawOption) *Image {
	ff := f.Copy()
	ff.drawOption = opt.drawOption
	ff.recalculateOption()
	return ff
}

func (f *Image) UpdateImageFromImage(img image.Image) {
	f.UpdateImage(ebiten.NewImageFromImage(img))
}

func (f *Image) UpdateImage(img *ebiten.Image) {
	f.img = img
	f.recalculateOption()
}

func (f *Image) recalculateOption() {
	bounds := f.img.Bounds()
	f.updateReference(float64(bounds.Dx()), float64(bounds.Dy()))
}

// Extension of DrawOption

func (f Image) WithMovement(x, y float64, replace ...bool) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withMovement(x, y, replace...)
	return ff
}

func (f Image) WithScaleRatio(x, y float64, replace ...bool) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withScaleRatio(x, y, replace...)
	return ff
}

func (f Image) WithRotation(degree float64, replace ...bool) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withRotation(degree, replace...)
	return ff
}

func (f Image) WithAlignment(a Align) *Image {
	ff := f.Copy()
	ff.drawOption = ff.drawOption.withAlignment(a)
	return ff
}
