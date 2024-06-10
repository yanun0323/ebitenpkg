package ebitenpkg

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type DebugDrawable interface {
	DebugDraw(screen *ebiten.Image, borderWidth ...int)
}
