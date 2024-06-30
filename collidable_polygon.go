package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type collidablePolygon struct {
	Collidable
	Controller

	debugImgCache *ebiten.Image
}

func NewCollidablePolygon(s Space, t BodyType, w, h float64, a ...Align) CollidablePolygon {
	ctr := NewController(w, h, a...)

	cp := &collidablePolygon{
		Collidable: newCollidable(s, t, ctr),
		Controller: ctr,
	}

	s.AddBody(cp)

	return cp
}

/*
	Drawable
*/

func (cp *collidablePolygon) Draw(screen *ebiten.Image) {}

func (cp *collidablePolygon) DebugDraw(screen *ebiten.Image, clr ...color.Color) {
	if cp.debugImgCache == nil {
		w, h := cp.bound()
		cp.debugImgCache = DebugImage(int(w), int(h), clr...)
	}

	screen.DrawImage(cp.debugImgCache, cp.DrawOption())
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
	cp.cleanCache()
	return cp
}

/*
	CollidablePolygon
*/

func (cp *collidablePolygon) NewImage(img sysimage.Image) Image {
	return newImage(img, cp.Controller)
}

func (cp *collidablePolygon) NewText(s string, size float64) Text {
	return newText(s, size, cp.Controller)
}

/*
	private
*/

func (cp *collidablePolygon) cleanCache() {
	cp.debugImgCache = nil
}
