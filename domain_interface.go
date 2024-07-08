package ebitenpkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image interface {
	Drawable
	Controllable[Image]

	Border(clr color.Color, width int) Image
	Copy() Image
	ReplaceImage(img *ebiten.Image) Image
	EbitenImage() *ebiten.Image
}

type Text interface {
	Drawable
	Controllable[Text]

	Copy() Text
	SetColor(clr color.Color) Text
	SetLineSpacing(ls float64) Text
	SetText(s string) Text
	SetSize(size float64) Text

	Color() color.Color
	LineSpacing() float64
	Text() string
	Size() float64
}

type Space interface {
	Update()
	AddBody(c Collidable) Space
	RemoveBody(id ID) Space
	IsCollided(id ID) bool
	GetCollided(id ID) []Collidable
}

type CollidableImage interface {
	Drawable
	Controllable[CollidableImage]
	Collidable

	Attach(parent Controllable[any]) CollidableImage
	Detach() CollidableImage
	Border(clr color.Color, width int) CollidableImage
	Copy() CollidableImage
	ReplaceImage(img *ebiten.Image) CollidableImage
	EbitenImage() *ebiten.Image
}

type CollidablePolygon interface {
	Drawable
	Controllable[CollidablePolygon]
	Collidable

	Attach(parent Controllable[any]) CollidablePolygon
	Detach() CollidablePolygon
	ReplaceSize(w, h float64) CollidablePolygon
	Copy() CollidablePolygon
}
