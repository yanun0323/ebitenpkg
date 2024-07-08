package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type collidablePolygon struct {
	w, h   float64
	ctr    controller
	parent Controllable[any]
	cd     collider
	space  Space
	img    debugCache
}

func NewCollidablePolygon(space Space, bt CollisionType, w, h float64, a ...Align) CollidablePolygon {
	obj := &collidablePolygon{
		w:      w,
		h:      h,
		ctr:    newController(a...),
		parent: nil,
		cd:     newCollider(bt),
		space:  space,
	}

	space.AddBody(obj)

	return obj
}

/*
	Drawable
*/

func (cp collidablePolygon) Draw(screen *ebiten.Image) {
	screen.DrawImage(cp.img.Image(int(cp.w), int(cp.h)), cp.DrawOption())
}

func (cp *collidablePolygon) DebugDraw(screen *ebiten.Image, clr ...color.Color) {
	if len(clr) == 0 && cp.IsCollided() {
		clr = append(clr, _collidedColor)
	}

	screen.DrawImage(cp.img.Image(int(cp.w), int(cp.h), clr...), cp.DrawOption())

	drawVertexesAndBarycenter(screen, cp.ctr, cp.Vertexes())
}

/*
	Controllable
*/

func (cp *collidablePolygon) Align(a Align) CollidablePolygon {
	cp.ctr.Align(a)
	return cp
}

func (cp *collidablePolygon) Move(x, y float64, replace ...bool) CollidablePolygon {
	cp.ctr.Move(x, y, replace...)
	return cp
}

func (cp *collidablePolygon) Rotate(degree float64, replace ...bool) CollidablePolygon {
	cp.ctr.Rotate(degree, replace...)
	return cp
}

func (cp *collidablePolygon) Scale(x, y float64, replace ...bool) CollidablePolygon {
	cp.ctr.Scale(x, y, replace...)
	return cp
}

func (cp collidablePolygon) Aligned() Align {
	return cp.ctr.Aligned()
}

func (cp collidablePolygon) Moved() (x, y float64) {
	return cp.ctr.Moved()
}

func (cp collidablePolygon) Rotated() float64 {
	return cp.ctr.Rotated()
}

func (cp collidablePolygon) Scaled() (x, y float64) {
	return cp.ctr.Scaled()
}

func (cp collidablePolygon) DrawOption() *ebiten.DrawImageOptions {
	if cp.parent == nil {
		return getDrawOption(cp.w, cp.h, cp.ctr)
	}

	pW, pH := cp.parent.Bounds()
	return getDrawOption(cp.w, cp.h, cp.ctr, parent{w: pW, h: pH, ctr: cp.parent})
}

func (cp collidablePolygon) Bounds() (w, h float64) {
	return cp.w, cp.h
}

func (cp collidablePolygon) Barycenter() (float64, float64) {
	x, y := cp.ctr.Moved()
	if cp.parent == nil {
		return x, y
	}

	pX, pY := cp.parent.Moved()
	return x + pX, y + pY
}

/*
	Collidable
*/

func (cp collidablePolygon) ID() ID {
	return cp.cd.id
}

func (cp collidablePolygon) Type() CollisionType {
	return cp.cd.bt
}

func (cp collidablePolygon) IsCollided() bool {
	return cp.space.IsCollided(cp.ID())
}

func (cp collidablePolygon) GetCollided() []Collidable {
	return cp.space.GetCollided(cp.ID())
}

func (cp collidablePolygon) IsInside(x, y float64) bool {
	return isInside(cp.Vertexes(), Vector{X: x, Y: y})
}

func (cp collidablePolygon) Vertexes() []Vector {
	if cp.parent == nil {
		return getVertexes(cp.w, cp.h, cp.ctr)
	}

	pW, pH := cp.parent.Bounds()
	return getVertexes(cp.w, cp.h, cp.ctr, parent{w: pW, h: pH, ctr: cp.parent})
}

/*
	CollidablePolygon
*/

func (cp *collidablePolygon) Attach(parent Controllable[any]) CollidablePolygon {
	cp.parent = parent
	return cp
}

func (cp *collidablePolygon) Detach() CollidablePolygon {
	cp.parent = nil
	return cp
}

func (cp *collidablePolygon) ReplaceSize(w, h float64) CollidablePolygon {
	cp.w = w
	cp.h = h
	return cp
}

func (cp collidablePolygon) Copy() CollidablePolygon {
	cp.ctr = cp.ctr.Copy()
	cp.cd = newCollider(cp.cd.Type())

	cp.space.AddBody(&cp)

	return &cp
}
