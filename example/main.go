package main

import (
	"fmt"
	"image"
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
	Count        int
	Space        ebitenpkg.Space
	Walls        []ebitenpkg.CollidablePolygon
	Player       ebitenpkg.CollidableImage
	PlayerWeapon ebitenpkg.CollidablePolygon
	Opponent     ebitenpkg.CollidableImage
	GameInfo     ebitenpkg.Text

	PikachuAnime        *ebiten.Image
	PikachuSpirit       *ebiten.Image
	PikachuAnimeResult  *ebiten.Image
	PikachuSpiritResult *ebiten.Image
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

	player := ebitenpkg.NewCollidableImage(space, TypePlayer, helper.GopherImage(), ebitenpkg.AlignCenter).Move(100, 100)
	weapon := ebitenpkg.NewCollidablePolygon(space, TypePlayer, 100, 30, ebitenpkg.AlignCenter).Move(20, 10).Rotate(-30).Attach(player)

	pikachuAnime := ebiten.NewImageFromImage(helper.PikachuAnimeImage())
	pikachuSpirit := ebiten.NewImageFromImage(helper.PikachuSpiritImage())

	return &Game{
		Space:        space,
		Player:       player,
		PlayerWeapon: weapon,
		Opponent:     ebitenpkg.NewCollidableImage(space, TypeOpponent, helper.GopherImage(), ebitenpkg.AlignTopCenter).Move(200, 200).Scale(-1, 1),
		Walls: []ebitenpkg.CollidablePolygon{
			ebitenpkg.NewCollidablePolygon(space, TypeWall, 10, fH, ebitenpkg.AlignTopCenter).Move(20, 0),
			ebitenpkg.NewCollidablePolygon(space, TypeWall, 10, fH, ebitenpkg.AlignTopCenter).Move(fW-20, 0),
		},
		GameInfo:            ebitenpkg.NewText("Hello, World!", 20, ebitenpkg.AlignTopLeading).Move(10, 0).SetColor(color.White),
		PikachuAnime:        pikachuAnime,
		PikachuSpirit:       pikachuSpirit,
		PikachuAnimeResult:  pikachuAnime,
		PikachuSpiritResult: pikachuSpirit,
	}
}

func (g *Game) Update() error {
	g.Count++
	g.Space.Update()
	g.GameInfo.SetText(fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	pressed := inpututil.AppendPressedKeys(nil)

	i := (g.Count / 5 /* second */) % 6
	w, h := g.PikachuAnime.Bounds().Dx(), g.PikachuAnime.Bounds().Dy()/6
	sx, sy := 0, h*i

	g.PikachuAnimeResult = g.PikachuAnime.SubImage(image.Rect(sx, sy, sx+w, sy+h)).(*ebiten.Image)

	helper.InputHandler[ebitenpkg.CollidableImage]{
		Object:         g.Player,
		MoveUp:         true,
		MoveDown:       true,
		MoveLeft:       true,
		MoveLeftScale:  true,
		MoveRight:      true,
		MoveRightScale: true,
		RotateLeft:     true,
		RotateRight:    true,
	}.Update(pressed)

	helper.InputHandler[ebitenpkg.CollidablePolygon]{
		Object:         g.PlayerWeapon,
		MoveLeftScale:  true,
		MoveRightScale: true,
	}.Update(pressed)

	runtime.GC()
	return nil
}

var _debugColor = color.RGBA{G: 255, A: 255}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Opponent.Draw(screen, _debugColor)
	g.Player.Draw(screen, _debugColor)
	g.PlayerWeapon.Draw(screen, _debugColor)

	if g.PikachuAnimeResult != nil {
		screen.DrawImage(g.PikachuAnimeResult, nil)
	}

	g.GameInfo.Draw(screen)
	for _, w := range g.Walls {
		w.Draw(screen, _debugColor)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
