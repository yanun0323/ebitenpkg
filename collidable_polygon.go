package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type collidablePolygon struct {
	Collidable
	Controller

	debugImg debugImage
}

func NewCollidablePolygon(s Space, t BodyType, w, h float64, a ...Align) CollidablePolygon {
	ctr := NewController(w, h, a...)

	cp := &collidablePolygon{
		Collidable: newCollidable(s, t, ctr),
		Controller: ctr,
		debugImg:   newDebugImage(ctr),
	}

	s.AddBody(cp)

	return cp
}

/*
	Drawable
*/

func (cp *collidablePolygon) Draw(screen *ebiten.Image) {}

func (cp *collidablePolygon) DebugDraw(screen *ebiten.Image, clr ...color.Color) {
	cp.debugImg.Draw(screen, clr)
}

/*
	embedController
*/

func (cp *collidablePolygon) Align(a Align) CollidablePolygon {
	cp.Controller.Align(a)
	return cp
}

func (cp *collidablePolygon) Move(x, y float64, replace ...bool) CollidablePolygon {
	cp.Controller.Move(x, y, replace...)
	return cp
}

func (cp *collidablePolygon) Rotate(degree float64, replace ...bool) CollidablePolygon {
	cp.Controller.Rotate(degree, replace...)
	return cp
}

func (cp *collidablePolygon) Scale(x, y float64, replace ...bool) CollidablePolygon {
	cp.Controller.Scale(x, y, replace...)
	return cp
}

func (cp *collidablePolygon) updateControllerReference() CollidablePolygon {
	cp.Controller = cp.Controller.updateControllerReference()
	cp.debugImg.CleanCache()
	return cp
}

/*
	private
*/
