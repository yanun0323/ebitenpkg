package ebitenpkg

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image interface {
	Draw(screen *ebiten.Image)

	Align(align Align) Image
	Move(x, y float64, replace ...bool) Image
	Moving(x, y float64, tick int, replace ...bool) Image
	Scale(x, y float64, replace ...bool) Image
	Scaling(x, y float64, tick int, replace ...bool) Image
	Rotate(angle float64, replace ...bool) Image
	Rotating(angle float64, tick int, replace ...bool) Image
	Spriteable(SpriteSheetOption) Image
	Attach(parent Attachable) Image
	Detach() (parent Attachable)
	Collidable(space Space, group int) Image
	Debug(on ...bool) Image

	Bounds() (width int, height int)
	Aligned() Align
	Moved() (x, y float64)
	Scaled() (x, y float64)
	Rotated() (angle float64)
	Debugged() bool
	SpriteSheet() (SpriteSheetOption, bool)

	CollisionID() ID
	CollisionGroup() int
	Parent() Attachable
	IsClick(x, y float64) bool
}

type SpriteSheetOption struct {
	SpriteWidthCount  int
	SpriteHeightCount int
	SpriteHandler     func(fps, timestamp int, direction Direction) (indexX, indexY, scaleX, scaleY int)
}

func NewImage(img image.Image) Image {
	return &eImage{
		id:          newID(),
		image:       ebiten.NewImageFromImage(img),
		imageWidth:  img.Bounds().Dx(),
		imageHeight: img.Bounds().Dy(),
	}
}

func NewImageWith(w, h int, clr ...color.Color) Image {
	return &eImage{
		id:          newID(),
		image:       NewEbitenImage(w, h, clr...),
		imageWidth:  w,
		imageHeight: h,
	}
}

type eImage struct {
	controller
	id ID

	image          *ebiten.Image
	imageWidth     int
	imageHeight    int
	draw           *ebiten.Image
	drawX, drawY   int
	drawSX, drawSY int

	spriteOption SpriteSheetOption

	parent Attachable

	collisionSpace Space
	collisionGroup int

	debug *ebiten.Image
}

func (e *eImage) Draw(screen *ebiten.Image) {
	if e.spriteOption.SpriteHandler == nil {
		option := getDrawOption(e.imageWidth, e.imageHeight, e.controller, 1, 1, e.parent)
		screen.DrawImage(e.image, option)

		if e.debug != nil {
			screen.DrawImage(e.debug, option)
		}
		return
	}

	sW, sH := e.Bounds()
	x, y, sX, sY := e.spriteOption.SpriteHandler(ebiten.DefaultTPS, CurrentGameTime(), e.controller.GetDirection())
	if x >= 0 && y >= 0 {
		if e.draw == nil || x != e.drawX || y != e.drawY || sX != e.drawSX || sY != e.drawSY {
			oX, oY := x*sW, y*sH

			rect := image.Rect(oX, oY, oX+sW, oY+sH)
			e.draw = e.image.SubImage(rect).(*ebiten.Image)
		}

		e.drawX, e.drawY = x, y
		e.drawSX, e.drawSY = sX, sY
	}

	option := getDrawOption(sW, sH, e.controller, float64(sX), float64(sY), e.parent)

	if e.draw != nil {
		screen.DrawImage(e.draw, option)
	}

	if e.debug != nil {
		screen.DrawImage(e.debug, option)
	}
}

func (e *eImage) Align(align Align) Image {
	e.controller.SetAlign(align)
	return e
}

func (e *eImage) Move(x, y float64, replace ...bool) Image {
	e.controller.SetMove(x, y, replace...)
	return e
}

func (e *eImage) Moving(x, y float64, tick int, replace ...bool) Image {
	e.controller.SetMoving(x, y, tick, replace...)
	return e
}

func (e *eImage) Scale(x, y float64, replace ...bool) Image {
	e.controller.SetScale(x, y, replace...)
	return e
}

func (e *eImage) Scaling(x, y float64, tick int, replace ...bool) Image {
	e.controller.SetScaling(x, y, tick, replace...)
	return e
}

func (e *eImage) Rotate(angle float64, replace ...bool) Image {
	e.controller.SetRotate(angle, replace...)
	return e
}

func (e *eImage) Rotating(angle float64, tick int, replace ...bool) Image {
	e.controller.SetRotating(angle, tick, replace...)
	return e
}

func (e *eImage) Spriteable(opt SpriteSheetOption) Image {
	e.spriteOption = opt
	return e
}

func (e *eImage) Attach(parent Attachable) Image {
	e.parent = parent
	return e
}

func (e *eImage) Detach() (parent Attachable) {
	result := e.parent
	e.parent = nil
	return result
}

func (e *eImage) Collidable(space Space, group int) Image {
	if e.collisionSpace != nil {
		e.collisionSpace.RemoveBody(e)
	}

	e.collisionSpace = space.AddBody(e)
	e.collisionGroup = group
	return e
}

func (e *eImage) Debug(on ...bool) Image {
	if len(on) != 0 && !on[0] {
		e.debug = nil
		return e
	}

	if e.debug != nil {
		return e
	}

	e.debug = NewEbitenImageWith(e.Bounds, DefaultDebugColor())
	return e
}

func (e *eImage) Bounds() (width int, height int) {
	if e.spriteOption.SpriteHandler == nil {
		return e.imageWidth, e.imageHeight
	}

	return e.imageWidth / e.spriteOption.SpriteWidthCount, e.imageHeight / e.spriteOption.SpriteHeightCount
}

func (e *eImage) Aligned() Align {
	return e.controller.GetAlign()
}

func (e *eImage) Moved() (x, y float64) {
	return e.controller.GetMove()
}

func (e *eImage) Scaled() (x, y float64) {
	return e.controller.GetScale()
}

func (e *eImage) Rotated() (angle float64) {
	return e.controller.GetRotate()
}

func (e *eImage) Debugged() bool {
	return e.debug != nil
}

func (e *eImage) SpriteSheet() (SpriteSheetOption, bool) {
	return e.spriteOption, e.spriteOption.SpriteHandler != nil
}

func (e *eImage) CollisionID() ID {
	return e.id
}

func (e *eImage) CollisionGroup() int {
	return e.collisionGroup
}

func (e *eImage) Parent() Attachable {
	return e.parent
}

func (e *eImage) IsClick(x, y float64) bool {
	vertexes := getVertexes(float64(e.imageWidth), float64(e.imageHeight), e, e.parent)
	return isInside(vertexes, Vector{X: x, Y: y})
}
