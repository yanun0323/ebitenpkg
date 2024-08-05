package ebitenpkg

import "github.com/hajimehoshi/ebiten/v2"

type Game[T any] struct {
	Data    T
	Updates func(T) error
	Draws   func(data T, screen *ebiten.Image)
	Layouts func(data T, outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

func (g Game[T]) Update() error {
	return g.Updates(g.Data)
}

func (g Game[T]) Draw(screen *ebiten.Image) {
	g.Draws(g.Data, screen)
}

func (g Game[T]) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Layouts(g.Data, outsideWidth, outsideHeight)
}

func (g Game[T]) Run() error {
	return ebiten.RunGame(g)
}
