package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	// Draw is an alias to screen.DrawImage(img.Image(), img.Option())
	Draw(screen *ebiten.Image)
	DebugDraw(screen *ebiten.Image, clr ...color.Color)
}

type Attachable interface {
	Aligned() Align
	Moved() (x, y float64)
	Rotated() float64
	Scaled() (x, y float64)
	Bounds() (w, h float64)
}

type Controllable[T any] interface {
	Align(a Align) T
	Move(x, y float64, replace ...bool) T
	Rotate(degree float64, replace ...bool) T
	Scale(x, y float64, replace ...bool) T
	Aligned() Align
	Moved() (x, y float64)
	Rotated() float64
	Scaled() (x, y float64)
	DrawOption() *ebiten.DrawImageOptions
	Bounds() (w, h float64)
	Barycenter() (x, y float64)
}

type Collidable interface {
	ID() ID
	Type() CollisionType
	// IsCollided should be call after Space.Update()
	IsCollided() bool
	// GetCollided should be call after Space.Update()
	GetCollided() []Collidable
	IsInside(x, y float64) bool
	// Barycenter() (x, y float64)
	Vertexes() []Vector
}
