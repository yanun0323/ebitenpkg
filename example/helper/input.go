package helper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type InputHandler[T ebitenpkg.Controllable[T]] struct {
	Object T
}

func (i *InputHandler[T]) Update(keys []ebiten.Key) {
	for _, in := range keys {
		switch in {
		case ebiten.KeyW:
			i.Object.Move(0, -5)
		case ebiten.KeyS:
			i.Object.Move(0, 5)
		case ebiten.KeyA:
			i.Object.Move(-5, 0)
			if scaleX, _ := i.Object.Scaled(); scaleX > 0 {
				i.Object.Rotate(-i.Object.Rotated(), true)
			}
			i.Object.Scale(-1, 1, true)
		case ebiten.KeyD:
			i.Object.Move(5, 0)
			if scaleX, _ := i.Object.Scaled(); scaleX < 0 {
				i.Object.Rotate(-i.Object.Rotated(), true)
			}
			i.Object.Scale(1, 1, true)
		case ebiten.KeyQ:
			i.Object.Rotate(-1)
		case ebiten.KeyE:
			i.Object.Rotate(1)
		}
	}
}
