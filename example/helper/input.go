package helper

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type InputHandler[T ebitenpkg.Controllable[T]] struct {
	Object         T
	MoveUp         bool
	MoveDown       bool
	MoveLeft       bool
	MoveLeftScale  bool
	MoveRight      bool
	MoveRightScale bool
	RotateLeft     bool
	RotateRight    bool
}

func (i InputHandler[T]) Update(keys []ebiten.Key) {
	for _, in := range keys {
		switch in {
		case ebiten.KeyW:
			if i.MoveUp {
				i.Object.Move(0, -5)
			}
		case ebiten.KeyS:
			if i.MoveDown {
				i.Object.Move(0, 5)
			}
		case ebiten.KeyA:
			if i.MoveLeft {
				i.Object.Move(-5, 0)
			}
			if i.MoveLeftScale {
				if scaleX, _ := i.Object.Scaled(); scaleX > 0 {
					i.Object.Rotate(-i.Object.Rotated(), true)
				}
				i.Object.Scale(-1, 1, true)
			}
		case ebiten.KeyD:
			if i.MoveRight {
				i.Object.Move(5, 0)
			}
			if i.MoveRightScale {
				if scaleX, _ := i.Object.Scaled(); scaleX < 0 {
					i.Object.Rotate(-i.Object.Rotated(), true)
				}
				i.Object.Scale(1, 1, true)
			}
		case ebiten.KeyQ:
			if i.RotateLeft {
				i.Object.Rotate(-1)
			}
		case ebiten.KeyE:
			if i.RotateRight {
				i.Object.Rotate(1)
			}
		}
	}
}
