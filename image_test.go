package ebitenpkg

import (
	"image"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestImage(t *testing.T) {
	var (
		oi     image.Image   = ebiten.NewImage(50, 50)
		screen *ebiten.Image = ebiten.NewImage(100, 100)
	)

	img := NewImage(oi).
		Align(AlignCenter).
		Move(0, 0).
		Scale(1, 1).
		Rotate(30).
		Spriteable(SpriteSheetOption{})

	img.Draw(screen)
}
