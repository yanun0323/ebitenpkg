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
	Detach()
	Collidable(space Space, group int) Image
	Debug(on ...bool) Image
	HandleImage(handler func(*ebiten.Image)) Image

	Bounds() (width int, height int)
	Aligned() Align
	Moved() (x, y float64)
	Scaled() (x, y float64)
	Rotated() (angle float64)
	Debugged() bool
	SpriteSheet() (SpriteSheetOption, bool)

	ID() ID
	Group() int
	Parent() Attachable
	IsClick(x, y float64) bool
}

type SpriteSheetOption struct {
	SpriteWidthCount  int
	SpriteHeightCount int
	SpriteHandler     func(fps, timestamp int, direction Direction) (indexX, indexY, scaleX, scaleY int)
}

func NewImage(img image.Image, children ...Attachable) Image {
	result := &eImage{
		id:          newValue(newID()),
		image:       newValue(ebiten.NewImageFromImage(img)),
		imageBounds: newValue(img.Bounds()),
	}

	for _, s := range children {
		attach(result, s)
	}

	return result
}

func NewRectangle(w, h int, clr color.Color, children ...Attachable) Image {
	result := &eImage{
		id:          newValue(newID()),
		image:       newValue(NewEbitenImage(w, h, clr)),
		imageBounds: newValue(image.Rect(0, 0, w, h)),
	}

	for _, s := range children {
		attach(result, s)
	}

	return result
}

type eImage struct {
	controller
	id value[ID]

	image       value[*ebiten.Image]
	imageBounds value[image.Rectangle]
	draw        value[*ebiten.Image]
	drawCoords  value[image.Point]
	drawScale   value[image.Point]

	spriteOption value[SpriteSheetOption]

	parent   value[Attachable]
	children slices[Attachable]

	collisionSpace value[Space]
	collisionGroup value[int]

	debug value[*ebiten.Image]
}

func (e *eImage) Draw(screen *ebiten.Image) {
	defer func() {
		e.children.Range(func(_ int, c Attachable) bool {
			c.Draw(screen)
			return true
		})
	}()

	spriteOption := e.spriteOption.Load()
	if spriteOption.SpriteHandler == nil {
		imageBounds := e.imageBounds.Load()
		option := getDrawOption(imageBounds.Dx(), imageBounds.Dy(), e.controller, 1, 1, e.Parent())
		screen.DrawImage(e.image.Load(), option)

		if debug := e.debug.Load(); debug != nil {
			screen.DrawImage(debug, option)
		}
		return
	}

	sW, sH := e.Bounds()
	x, y, sX, sY := spriteOption.SpriteHandler(ebiten.DefaultTPS, CurrentGameTime(), e.controller.GetDirection())
	if x >= 0 && y >= 0 {
		draw := e.draw.Load()
		drawCoords := e.drawCoords.Load()
		drawScale := e.drawScale.Load()
		if draw == nil || x != drawCoords.X || y != drawCoords.Y || sX != drawScale.X || sY != drawScale.Y {
			oX, oY := x*sW, y*sH

			rect := image.Rect(oX, oY, oX+sW, oY+sH)
			draw = e.image.Load().SubImage(rect).(*ebiten.Image)
			e.draw.Store(draw)
		}

		e.drawCoords.Store(image.Point{X: x, Y: y})
		e.drawScale.Store(image.Point{X: sX, Y: sY})
	}

	option := getDrawOption(sW, sH, e.controller, float64(sX), float64(sY), e.Parent())

	if draw := e.draw.Load(); draw != nil {
		screen.DrawImage(draw, option)
	}

	if debug := e.debug.Load(); debug != nil {
		screen.DrawImage(debug, option)
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
	e.spriteOption.Store(opt)
	return e
}

func (e *eImage) Attach(parent Attachable) Image {
	attach(parent, e)
	return e
}

func (e *eImage) Detach() {
	detach(e)
}

func (e *eImage) Collidable(space Space, group int) Image {
	sp := e.collisionSpace.Load()
	if sp != nil {
		sp.RemoveBody(e)
	}

	e.collisionSpace.Store(space.AddBody(e))
	e.collisionGroup.Store(group)
	return e
}

func (e *eImage) Debug(on ...bool) Image {
	if len(on) != 0 && !on[0] {
		e.debug.Delete()
		return e
	}

	if e.debug.Load() != nil {
		return e
	}

	e.debug.Store(NewEbitenImageFromBounds(e.Bounds, DefaultDebugColor()))
	return e
}

func (e *eImage) HandleImage(handler func(*ebiten.Image)) Image {
	handler(e.image.Load())
	return e
}

func (e *eImage) Bounds() (width int, height int) {
	imageBounds := e.imageBounds.Load()
	opt := e.spriteOption.Load()
	if opt.SpriteHandler == nil {
		return imageBounds.Dx(), imageBounds.Dy()
	}

	return imageBounds.Dx() / opt.SpriteWidthCount, imageBounds.Dy() / opt.SpriteHeightCount
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
	return e.debug.Load() != nil
}

func (e *eImage) SpriteSheet() (SpriteSheetOption, bool) {
	opt := e.spriteOption.Load()
	return opt, opt.SpriteHandler != nil
}

func (e *eImage) ID() ID {
	return e.id.Load()
}

func (e *eImage) Group() int {
	return e.collisionGroup.Load()
}

func (e *eImage) Parent() Attachable {
	return e.parent.Load()
}

func (e *eImage) IsClick(x, y float64) bool {
	imageBounds := e.imageBounds.Load()
	vertexes := getVertexes(float64(imageBounds.Dx()), float64(imageBounds.Dy()), e, e.Parent())
	return isInside(vertexes, Vector{X: x, Y: y})
}
