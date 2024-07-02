package ebitenpkg

import (
	sysimage "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	// Draw is an alias to screen.DrawImage(img.Image(), img.Option())
	Draw(screen *ebiten.Image)
	DebugDraw(screen *ebiten.Image, clr ...color.Color)
}

//go:generate domaingen -destination=controller.go -name=controller -package=ebitenpkg
type Controller interface {
	embedController[Controller]

	Copy() Controller
	NewImage(sysimage.Image) Image
	NewText(s string, size float64) Text
}

//go:generate domaingen -destination=image.go -name=image -package=ebitenpkg -noembed
type Image interface {
	Drawable
	embedController[Image]

	Border(color.Color, int) Image
	Copy(with ...Controller) Image
	GetController() Controller
	ReplaceImage(*ebiten.Image) Image
	EbitenImage() *ebiten.Image
	Vertexes() []Vector
}

//go:generate domaingen -destination=text.go -name=text -package=ebitenpkg -noembed
type Text interface {
	Drawable
	embedController[Text]

	Copy(with ...Controller) Text
	SetColor(color.Color) Text
	SetLineSpacing(float64) Text
	SetText(string) Text
	SetSize(float64) Text

	Bound() (w, h float64)
	Color() color.Color
	GetController() Controller
	LineSpacing() float64
	Text() string
	Size() float64
	Vertexes() []Vector
}

//go:generate domaingen -destination=space.go -name=space -package=ebitenpkg
type Space interface {
	Update()
	AddBody(Collidable) Space
	RemoveBody(ID) Space
	IsCollided(ID) bool
	GetCollided(ID) []Collidable
}

//go:generate domaingen -destination=collidable_image.go -name=collidableImage -package=ebitenpkg -noembed
type CollidableImage interface {
	Drawable
	Collidable
	embedController[CollidableImage]

	GetImage() Image
}

//go:generate domaingen -destination=collidable_polygon.go -name=collidablePolygon -package=ebitenpkg -noembed
type CollidablePolygon interface {
	Drawable
	Collidable
	embedController[CollidablePolygon]
}

//go:generate domaingen -destination=body.go -name=body -package=ebitenpkg -noembed
type Collidable interface {
	ID() ID
	Type() BodyType

	// IsCollided should be call after Space.Update()
	IsCollided() bool
	IsCollide(Vector) bool

	// GetCollided should be call after Space.Update()
	GetCollided() []Collidable

	CollideCenter() Vector

	controller() Controller
}

type embedController[T any] interface {
	Align(Align) T
	Move(x, y float64, replace ...bool) T
	Rotate(degree float64, replace ...bool) T
	Scale(x, y float64, replace ...bool) T
	updateControllerReference() T

	Aligned() Align
	Moved() (x, y float64)
	Rotated() float64
	Scaled() (x, y float64)

	DrawOption() *ebiten.DrawImageOptions

	updateReference(x, y float64)
	rotationCenter() Vector
	vertexes() []Vector
	bound() (w, h float64)
}
