package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

func main() {
	g, err := newGame()
	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

var _ ebiten.Game = (*game)(nil)

func newGame() (ebiten.Game, error) {
	f, err := os.Open("./example/go.png")
	if err != nil {
		return nil, fmt.Errorf("read file go.png, err: %w", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decode png, err: %w", err)
	}

	g := &game{
		img1: ebitenpkg.NewImageFromImage(img).WithMovement(100, 0),
		tps:  ebitenpkg.NewText("", 30).WithAlignment(ebitenpkg.AlignTrailing).WithFontColor(color.White).WithFontLinsSpace(3),
		fps:  ebitenpkg.NewText("", 30).WithAlignment(ebitenpkg.AlignTopLeading).WithFontColor(color.White).WithFontLinsSpace(3),
		opt:  ebitenpkg.NewDrawOption(0, 0),
	}

	g.img1.Move(500, 100)
	g.img1.Scale(-1, -1)

	tw, th := g.tps.Size()
	g.tps.Move(tw, th)

	g.opt.Move(300, 200)

	return g, nil
}

type game struct {
	img1     *ebitenpkg.Image
	tps, fps *ebitenpkg.Text
	opt      *ebitenpkg.DrawOption
}

func (g *game) Update() error {
	g.img1.Rotate(1)

	g.tps.Rotate(-1)
	g.tps.SetText(fmt.Sprintf("TPS: %d", ebiten.TPS()))

	g.fps.SetText(fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()))

	g.opt.Move(0.3, 0.3)
	g.opt.Rotate(-45, true)
	g.opt.Scale(1.001, 1.001)

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.img1.DebugDraw(screen)
	g.tps.Draw(screen)
	g.fps.DebugDraw(screen)
	g.opt.Text("OPTION", 40).Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1366, 768
}
