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
	Color(r, g, b, a uint8) Image
	Coloring(r, g, b, a uint8, tick int) Image
	Spriteable(SpriteSheetOption) Image
	Attach(parent Attachable) Image
	Detach()
	Collidable(space Space, group int) Image
	Debug(on ...bool) Image
	HandleImage(handler func(*ebiten.Image)) Image

	WithAnimation(animation Animation) Image
	Animation() Animation

	Bounds() (width int, height int)
	Aligned() Align
	Moved() (x, y float64)
	Scaled() (x, y float64)
	Rotated() (angle float64)
	Colored() (r, g, b, a uint8)
	Debugged() bool
	SpriteSheet() (SpriteSheetOption, bool)

	ID() ID
	Group() int
	Parent() Attachable
	IsClick(x, y float64) bool
}

type SpriteSheetOption struct {
	SpriteColumnCount int
	SpriteRowCount    int
	SpriteIndexCount  int
	SpriteHandler     func(fps, timestamp int, direction Direction) (index int, scaleX, scaleY int)
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

func NewRoundedRectangle(w, h, round int, clr color.Color, children ...Attachable) Image {
	result := &eImage{
		id:          newValue(newID()),
		image:       newValue(NewEbitenRoundedImage(w, h, round, clr)),
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

func (e *eImage) drawChildren(screen *ebiten.Image) {
	e.children.Range(func(_ int, c Attachable) bool {
		c.Draw(screen)
		return true
	})
}

func (e *eImage) getDrawOption(w int, h int, current *controller, tempScaleX float64, tempScaleY float64, pr ...Attachable) *ebiten.DrawImageOptions {
	opt := getDrawOption(w, h, current, tempScaleX, tempScaleY, pr...)
	r, g, b, a := e.GetColor()
	opt.ColorScale.Reset()
	opt.ColorScale.Scale(
		float32(r)/255,
		float32(g)/255,
		float32(b)/255,
		1,
	)
	opt.ColorScale.ScaleAlpha(float32(a) / 255)
	return opt
}

func (e *eImage) Draw(screen *ebiten.Image) {
	spriteOption := e.spriteOption.Load()
	if spriteOption.SpriteHandler == nil {
		imageBounds := e.imageBounds.Load()
		option := e.getDrawOption(imageBounds.Dx(), imageBounds.Dy(), &e.controller, 1, 1, e.Parent())
		screen.DrawImage(e.image.Load(), option)

		if debug := e.debug.Load(); debug != nil {
			screen.DrawImage(debug, option)
		}

		e.drawChildren(screen)

		return
	}

	sW, sH := e.Bounds()
	idx, sX, sY := spriteOption.SpriteHandler(ebiten.DefaultTPS, CurrentGameTime(), e.controller.GetDirection())
	idx = idx % spriteOption.SpriteIndexCount
	x := idx % spriteOption.SpriteColumnCount
	y := idx / spriteOption.SpriteColumnCount
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

	option := e.getDrawOption(sW, sH, &e.controller, float64(sX), float64(sY), e.Parent())

	if draw := e.draw.Load(); draw != nil {
		screen.DrawImage(draw, option)
	}

	if debug := e.debug.Load(); debug != nil {
		screen.DrawImage(debug, option)
	}

	e.drawChildren(screen)
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

func (e *eImage) Color(r, g, b, a uint8) Image {
	e.controller.SetColor(r, g, b, a)
	return e
}

func (e *eImage) Coloring(r, g, b, a uint8, tick int) Image {
	e.controller.SetColoring(r, g, b, a, tick)
	return e
}

func (e *eImage) Spriteable(opt SpriteSheetOption) Image {
	if opt.SpriteColumnCount == 0 {
		opt.SpriteColumnCount = 1
	}

	if opt.SpriteRowCount == 0 {
		opt.SpriteRowCount = 1
	}

	if opt.SpriteIndexCount == 0 {
		opt.SpriteIndexCount = 1
	}

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

func (e *eImage) WithAnimation(animation Animation) Image {
	e.controller.SetAnimation(animation)
	return e
}

func (e *eImage) Animation() Animation {
	return e.GetAnimation()
}

func (e *eImage) Bounds() (width int, height int) {
	imageBounds := e.imageBounds.Load()
	opt := e.spriteOption.Load()
	if opt.SpriteHandler == nil {
		return imageBounds.Dx(), imageBounds.Dy()
	}

	return imageBounds.Dx() / opt.SpriteColumnCount, imageBounds.Dy() / opt.SpriteRowCount
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

func (e *eImage) Colored() (r, g, b, a uint8) {
	return e.controller.GetColor()
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
