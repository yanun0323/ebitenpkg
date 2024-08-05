package ebitenpkg

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image interface {
	Draw(screen *ebiten.Image)

	Align(align Align) Image
	Move(x, y float64, replace ...bool) Image
	Scale(x, y float64, replace ...bool) Image
	Rotate(angle float64, replace ...bool) Image
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
	IsCollidable() bool
	CollisionID() ID
	CollisionGroup() int
}

type SpriteSheetOption struct {
	SpriteWidthCount  int
	SpriteHeightCount int
	SpriteHandler     func(fps, timestamp int, direction Direction) (indexX, indexY int)
}

func NewImage(img image.Image) Image {
	imageWidth, imageHeight := img.Bounds().Dx(), img.Bounds().Dy()
	return &eImage{
		id:          newID(),
		image:       ebiten.NewImageFromImage(img),
		imageWidth:  imageWidth,
		imageHeight: imageHeight,
	}
}

type eImage struct {
	controller
	id ID

	image        *ebiten.Image
	imageWidth   int
	imageHeight  int
	draw         *ebiten.Image
	drawX, drawY int

	spriteOption SpriteSheetOption

	parent Attachable

	collisionSpace Space
	collisionGroup int

	debug *ebiten.Image
}

func (e *eImage) Draw(screen *ebiten.Image) {

	if e.spriteOption.SpriteHandler == nil {
		option := getDrawOption(e.imageWidth, e.imageHeight, e.controller, e.parent)
		screen.DrawImage(e.image, option)

		if e.debug != nil {
			screen.DrawImage(e.debug, option)
		}
		return
	}

	sW, sH := e.Bounds()
	option := getDrawOption(sW, sH, e.controller, e.parent)
	x, y := e.spriteOption.SpriteHandler(ebiten.DefaultTPS, CurrentGameTime(), e.controller.GetDirection())
	if x >= 0 && y >= 0 && (e.draw == nil || x != e.drawX || y != e.drawY) {
		fmt.Printf("sW: %d, sH: %d\n", sW, sH)
		oX, oY := x*sW, y*sH+e.drawY

		rect := image.Rect(oX, oY, oX+sW, oY+sH)
		fmt.Printf("%+v\n", rect)
		e.draw = e.image.SubImage(rect).(*ebiten.Image)
	}

	if e.draw != nil {
		println("draw draw")
		screen.DrawImage(e.draw, option)
	}

	if e.debug != nil {
		println("draw debug")
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

func (e *eImage) Scale(x, y float64, replace ...bool) Image {
	e.controller.SetScale(x, y, replace...)
	return e
}

func (e *eImage) Rotate(angle float64, replace ...bool) Image {
	e.controller.SetRotate(angle, replace...)
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
		e.collisionSpace.RemoveBody(e.CollisionID())
	}

	e.collisionSpace = space.AddBody(e)
	e.collisionGroup = group
	return e
}

func (e *eImage) Debug(on ...bool) Image {
	if len(on) == 0 || !on[0] {
		e.debug = nil
		return e
	}

	w, h := e.Bounds()
	img := ebiten.NewImage(w, h)
	img.Fill(DefaultDebugColor())
	e.debug = img
	return e
}

func (e eImage) Bounds() (width int, height int) {
	if e.spriteOption.SpriteHandler == nil {
		return e.imageWidth, e.imageHeight
	}

	return e.imageWidth / e.spriteOption.SpriteWidthCount, e.imageHeight / e.spriteOption.SpriteHeightCount
}

func (e eImage) Aligned() Align {
	return e.controller.GetAlign()
}

func (e eImage) Moved() (x, y float64) {
	return e.controller.GetMove()
}

func (e eImage) Scaled() (x, y float64) {
	return e.controller.GetScale()
}

func (e eImage) Rotated() (angle float64) {
	return e.controller.GetRotate()
}

func (e eImage) Debugged() bool {
	return e.debug != nil
}

func (e eImage) SpriteSheet() (SpriteSheetOption, bool) {
	return e.spriteOption, e.spriteOption.SpriteHandler != nil
}

func (e eImage) IsCollidable() bool {
	return e.collisionSpace != nil
}

func (e eImage) CollisionID() ID {
	return e.id
}

func (e eImage) CollisionGroup() int {
	return e.collisionGroup
}

func (e eImage) Parent() Attachable {
	return e.parent
}
