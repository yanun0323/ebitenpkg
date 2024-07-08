package main

import (
	"fmt"
	"image/color"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yanun0323/ebitenpkg"
	"github.com/yanun0323/ebitenpkg/example/helper"
)

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}

type Game struct {
	Space    ebitenpkg.Space
	Walls    []ebitenpkg.CollidablePolygon
	Player   ebitenpkg.CollidableImage
	Opponent ebitenpkg.CollidableImage
	GameInfo ebitenpkg.Text
}

const (
	TypeWall ebitenpkg.CollisionType = iota
	TypePlayer
	TypeOpponent
)

func NewGame() ebiten.Game {
	space := ebitenpkg.NewSpace()
	w, h := ebiten.WindowSize()
	fW, fH := float64(w), float64(h)
	return &Game{
		Space:    space,
		Player:   ebitenpkg.NewCollidableImage(space, TypePlayer, helper.GopherImage(), ebitenpkg.AlignBottomCenter).Move(100, 100),
		Opponent: ebitenpkg.NewCollidableImage(space, TypeOpponent, helper.GopherImage(), ebitenpkg.AlignTopCenter).Move(200, 200).Scale(-1, 1),
		Walls: []ebitenpkg.CollidablePolygon{
			ebitenpkg.NewCollidablePolygon(space, TypeWall, 10, fH, ebitenpkg.AlignTopCenter).Move(20, 0),
			ebitenpkg.NewCollidablePolygon(space, TypeWall, 10, fH, ebitenpkg.AlignTopCenter).Move(fW-20, 0),
		},
		GameInfo: ebitenpkg.NewText("Hello, World!", 20, ebitenpkg.AlignTopLeading).Move(10, 0).SetColor(color.White),
	}
}

func (g *Game) Update() error {
	g.Space.Update()
	g.GameInfo.SetText(fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))

	player := helper.InputHandler[ebitenpkg.CollidableImage]{Object: g.Player}
	player.Update(inpututil.AppendPressedKeys(nil))

	runtime.GC()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Opponent.DebugDraw(screen)
	g.Player.DebugDraw(screen)

	g.GameInfo.Draw(screen)
	for _, w := range g.Walls {
		w.DebugDraw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
