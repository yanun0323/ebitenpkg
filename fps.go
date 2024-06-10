package ebitenpkg

import "github.com/hajimehoshi/ebiten/v2"

func InitFPS() {
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.TPS())
}
