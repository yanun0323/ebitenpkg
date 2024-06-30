package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"runtime"

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

type game struct {
	space    ebitenpkg.Space
	img1     ebitenpkg.CollidableImage
	tps, fps ebitenpkg.Text
	ctr      ebitenpkg.Controller
	grid     []ebitenpkg.Image
	center   ebitenpkg.CollidableImage
	rect     ebitenpkg.Image
	vertexes []ebitenpkg.Image
}

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
	s := ebitenpkg.NewSpace()
	img1 := ebitenpkg.NewCollidableImage(s, 0, img)

	gridSize := 50
	c := color.Gray{50}
	grid := make([]ebitenpkg.Image, 0, 400)
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			img := ebitenpkg.NewImage(ebiten.NewImage(gridSize, gridSize), ebitenpkg.AlignTopLeading).
				Border(c, 1).
				Move(float64(i*gridSize), float64(j*gridSize))
			grid = append(grid, img)
		}
	}

	rectSize := gridSize * 2
	rotate := 45.0

	centerColor := color.RGBA{50, 0, 0, 100}
	centerImg := ebiten.NewImage(rectSize, rectSize)
	centerImg.Fill(centerColor)

	rectColor := color.RGBA{255, 0, 0, 255}

	mX, mY := 8, 6
	a := ebitenpkg.AlignBottomTrailing
	rect := ebitenpkg.NewImage(ebiten.NewImage(rectSize, rectSize), a).
		Border(rectColor, 1).
		Move(float64(gridSize*mX), float64(gridSize*mY)).
		Scale(3, 2).
		Rotate(rotate)

	center := ebitenpkg.NewCollidableImage(s, 1, centerImg, a).
		Move(float64(gridSize*mX), float64(gridSize*mY)).
		Scale(1, 1).
		Rotate(rotate)

	g := &game{
		space:  s,
		img1:   img1,
		tps:    ebitenpkg.NewText("", 30).Align(ebitenpkg.AlignTrailing).SetColor(color.White).SetLineSpacing(3),
		fps:    ebitenpkg.NewText("", 30).Align(ebitenpkg.AlignTopLeading).SetColor(color.White).SetLineSpacing(3),
		ctr:    ebitenpkg.NewController(0, 0),
		grid:   grid,
		rect:   rect,
		center: center,
	}

	g.img1.
		Move(400, 50).
		Scale(-1, -1)

	tw, th := g.tps.Bound()
	g.tps.Move(tw, th)

	g.ctr.Move(300, 200)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return g, nil
}

func (g *game) Update() error {

	if g.img1.IsCollided() {
		println("GO Collided")
	} else {
		g.img1.Rotate(1)
	}

	g.tps.Rotate(-1)
	g.tps.SetText(fmt.Sprintf("TPS: %d", ebiten.TPS()))

	g.fps.SetText(fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()))

	g.ctr.Rotate(-45, true)

	if x, y := g.ctr.Scaled(); x < 2 && y < 2 {
		g.ctr.Move(0.3, 0.3)
		g.ctr.Scale(1.001, 1.001)
	}

	vs := g.rect.Vertexes()
	vs = append(vs, g.img1.GetImage().Vertexes()...)
	vs = append(vs, g.center.GetImage().Vertexes()...)
	vertexes := make([]ebitenpkg.Image, 0, len(vs))
	for _, v := range vs {
		img := ebitenpkg.NewImage(ebiten.NewImage(5, 5)).Move(v.X, v.Y).Border(color.White, 5)
		vertexes = append(vertexes, img)
	}
	g.vertexes = vertexes

	if err := g.space.Update(); err != nil {
		return err
	}

	// println()
	// for _, vv := range g.img1.GetImage().Vertexes() {
	// 	print(fmt.Sprintf("{%.2f, %.2f} ", vv.X, vv.Y))
	// }
	// println()

	runtime.GC()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for _, gr := range g.grid {
		gr.Draw(screen)
	}
	g.img1.DebugDraw(screen)
	g.tps.DebugDraw(screen)
	g.fps.DebugDraw(screen)
	g.ctr.NewText("OPTION", 40).DebugDraw(screen)
	g.center.Draw(screen)
	g.rect.Draw(screen)

	for _, v := range g.vertexes {
		v.Draw(screen)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
