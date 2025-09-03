package ebitenpkg

import "github.com/hajimehoshi/ebiten/v2"

type emptyImage struct{}

func EmptyImage() Image {
	return emptyImage{}
}

func (emptyImage) Draw(screen *ebiten.Image) {}
func (emptyImage) Detach()                   {}

func (ei emptyImage) Align(align Align) Image {
	return ei
}

func (ei emptyImage) Move(x, y float64, replace ...bool) Image {
	return ei
}

func (ei emptyImage) Moving(x, y float64, tick int, replace ...bool) Image {
	return ei
}

func (ei emptyImage) Scale(x, y float64, replace ...bool) Image {
	return ei
}

func (ei emptyImage) Scaling(x, y float64, tick int, replace ...bool) Image {
	return ei
}

func (ei emptyImage) Rotate(angle float64, replace ...bool) Image {
	return ei
}

func (ei emptyImage) Rotating(angle float64, tick int, replace ...bool) Image {
	return ei
}

func (ei emptyImage) Color(r, g, b, a uint8) Image {
	return ei
}
func (ei emptyImage) Coloring(r, g, b, a uint8, tick int) Image {
	return ei
}

func (ei emptyImage) Mask(x, y, w, h float64) Image {
	return ei
}

func (ei emptyImage) Masking(x, y, w, h float64, tick int) Image {
	return ei
}

func (ei emptyImage) Spriteable(SpriteSheetOption) Image {
	return ei
}

func (ei emptyImage) Attach(parent Attachable) Image {
	return ei
}

func (ei emptyImage) Collidable(space Space, group int) Image {
	return ei
}
func (ei emptyImage) Debug(on ...bool) Image {
	return ei
}
func (ei emptyImage) HandleImage(handler func(*ebiten.Image)) Image {
	return ei
}
func (ei emptyImage) WithAnimation(animation Animation) Image {
	return ei
}

func (emptyImage) Animation() Animation {
	return AnimationDefault()
}

func (emptyImage) Bounds() (width int, height int) {
	return 0, 0
}

func (emptyImage) Aligned() Align {
	return AlignCenter
}

func (emptyImage) Moved() (x, y float64) {
	return 0, 0
}

func (emptyImage) Scaled() (x, y float64) {
	return 0, 0
}

func (emptyImage) Rotated() (angle float64) {
	return 0
}

func (emptyImage) Colored() (r, g, b, a uint8) {
	return 255, 255, 255, 255
}

func (emptyImage) Masked() (x, y, w, h float64) {
	return 0, 0, 1, 1
}

func (emptyImage) Debugged() bool {
	return false
}

func (emptyImage) ID() ID {
	return ID{}
}

func (emptyImage) Group() int {
	return 0
}

func (emptyImage) Parent() Attachable {
	return nil
}

func (emptyImage) IsClick(x, y float64) bool {
	return false
}

func (ei emptyImage) Copy() Image {
	return ei
}
