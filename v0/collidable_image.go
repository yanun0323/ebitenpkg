package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type collidableImage struct {
	image
	parent Attachable
	cd     collider
	space  Space
}

func NewCollidableImage(space Space, bt CollisionType, img sysimage.Image, a Align, opt ...SpriteSheetOption) CollidableImage {
	var o SpriteSheetOption
	if len(opt) != 0 {
		o = opt[0]
	}

	obj := &collidableImage{
		image: image{
			ctr: newController(a),
			img: ebiten.NewImageFromImage(img),
			opt: o,
		},
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

func (ci collidableImage) Draw(screen *ebiten.Image, debug ...color.Color) {
	if len(debug) != 0 && ci.IsCollided() {
		debug[0] = _collidedColor
	}

	ci.image.Draw(screen, debug...)
}

/*
	Controllable
*/

func (ci *collidableImage) Align(a Align) CollidableImage {
	ci.ctr.Align(a)
	return ci
}

func (ci *collidableImage) Move(x, y float64, replace ...bool) CollidableImage {
	ci.ctr.Move(x, y, replace...)
	return ci
}

func (ci *collidableImage) Rotate(degree float64, replace ...bool) CollidableImage {
	ci.ctr.Rotate(degree, replace...)
	return ci
}

func (ci *collidableImage) Scale(x, y float64, replace ...bool) CollidableImage {
	ci.ctr.Scale(x, y, replace...)
	return ci
}

func (ci collidableImage) DrawOption() *ebiten.DrawImageOptions {
	w, h := ci.Bounds()
	if ci.parent == nil {
		return getDrawOption(w, h, ci.ctr)
	}

	pW, pH := ci.parent.Bounds()
	return getDrawOption(w, h, ci.ctr, parent{w: pW, h: pH, ctr: ci.parent})
}

func (ci collidableImage) Bounds() (w, h float64) {
	b := ci.img.Bounds()
	return float64(b.Dx()), float64(b.Dy())
}

/*
	Collidable
*/

func (ci collidableImage) ID() ID {
	return ci.cd.id
}

func (ci collidableImage) Type() CollisionType {
	return ci.cd.bt
}

func (ci collidableImage) IsCollided() bool {
	return ci.space.IsCollided(ci.ID())
}

func (ci collidableImage) GetCollided() []Collidable {
	return ci.space.GetCollided(ci.ID())
}

func (ci collidableImage) IsInside(x, y float64) bool {
	return isInside(ci.Vertexes(), Vector{X: x, Y: y})
}

func (ci collidableImage) Vertexes() []Vector {
	w, h := ci.Bounds()
	if ci.parent == nil {
		return getVertexes(w, h, ci.ctr)
	}

	pW, pH := ci.parent.Bounds()
	return getVertexes(w, h, ci.ctr, parent{w: pW, h: pH, ctr: ci.parent})
}

/*
	CollidableImage
*/

func (ci *collidableImage) Attach(parent Attachable) CollidableImage {
	ci.parent = parent
	return ci
}

func (ci *collidableImage) Detach() CollidableImage {
	ci.parent = nil
	return ci
}

func (ci *collidableImage) Border(clr color.Color, width int) CollidableImage {
	ci.image.Border(clr, width)
	return ci
}

func (ci collidableImage) Copy() CollidableImage {
	ci.image = ci.image.copy()
	ci.cd = newCollider(ci.cd.Type())

	ci.space.AddBody(&ci)

	return &ci
}

func (ci *collidableImage) ReplaceImage(img *ebiten.Image) CollidableImage {
	ci.image.ReplaceImage(img)
	return ci
}

func (ci collidableImage) EbitenImage() *ebiten.Image {
	return ci.image.EbitenImage()
}
