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
	Walls        []ebitenpkg.Image
	Player       ebitenpkg.Image
	PlayerWeapon ebitenpkg.Image
	Opponent     ebitenpkg.Image
	GameInfo     ebitenpkg.Text

	PikachuAnime        *ebiten.Image
	PikachuSprite       *ebiten.Image
	PikachuAnimeResult  *ebiten.Image
	PikachuSpriteResult *ebiten.Image

	pikachuAnimeImg ebitenpkg.Image
}

const (
	TypeWall int = iota
	TypePlayer
	TypeOpponent
)

func NewGame() ebiten.Game {
	space := ebitenpkg.NewSpace()
	w, h := ebiten.WindowSize()
	fW, fH := float64(w), float64(h)

	player := ebitenpkg.NewImage(helper.GopherImage()).
		Align(ebitenpkg.AlignCenter).
		Move(100, 100).
		Collidable(space, TypePlayer).
		Debug(true)

	weapon := ebitenpkg.NewImage(ebiten.NewImage(100, 30)).
		Align(ebitenpkg.AlignCenter).
		Move(20, 10).Rotate(-30).Attach(player).
		Collidable(space, TypePlayer).
		Debug(true)

	pikachuAnime := ebiten.NewImageFromImage(helper.PikachuAnimeImage())

	pikachuAnimeImg := ebitenpkg.NewImage(helper.PikachuAnimeImage()).
		Align(ebitenpkg.AlignTopLeading).
		Move(100, 100).
		Scale(1, 1).
		Spriteable(ebitenpkg.SpriteSheetOption{
			SpriteWidthCount:  1,
			SpriteHeightCount: 6,
			SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX, indexY int) {
				return 0, (timestamp / 5) % 6
			},
		}).
		Debug()

	pikachuSprite := ebiten.NewImageFromImage(helper.PikachuSpriteImage())

	return &Game{
		Space:        space,
		Player:       player,
		PlayerWeapon: weapon,
		Opponent:     ebitenpkg.NewImage(helper.GopherImage()).Align(ebitenpkg.AlignTop).Move(200, 200).Scale(-1, 1).Collidable(space, TypeOpponent).Debug(true),
		Walls: []ebitenpkg.Image{
			ebitenpkg.NewImage(ebiten.NewImage(10, int(fH))).Align(ebitenpkg.AlignTop).Move(20, 0).Collidable(space, TypeWall).Debug(true),
			ebitenpkg.NewImage(ebiten.NewImage(10, int(fH))).Align(ebitenpkg.AlignTop).Move(fW-20, 0).Collidable(space, TypeWall).Debug(true),
		},
		GameInfo:            ebitenpkg.NewText("Hello, World!", 20).Align(ebitenpkg.AlignTopLeading).Move(10, 0).SetColor(color.White).Debug(true),
		PikachuAnime:        pikachuAnime,
		pikachuAnimeImg:     pikachuAnimeImg,
		PikachuSprite:       pikachuSprite,
		PikachuAnimeResult:  pikachuAnime,
		PikachuSpriteResult: pikachuSprite,
	}
}

func (g *Game) Update() error {
	g.Count++
	ebitenpkg.GameUpdate()
	g.Space.GameUpdate()
	g.GameInfo.SetText(fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	pressed := inpututil.AppendPressedKeys(nil)

	i := (g.Count / 5 /* second */) % 6
	w, h := g.PikachuAnime.Bounds().Dx(), g.PikachuAnime.Bounds().Dy()/6
	sx, sy := 0, h*i

	g.PikachuAnimeResult = g.PikachuAnime.SubImage(image.Rect(sx, sy, sx+w, sy+h)).(*ebiten.Image)

	helper.InputHandler[ebitenpkg.Image]{
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

	helper.InputHandler[ebitenpkg.Image]{
		Object:         g.PlayerWeapon,
		MoveLeftScale:  true,
		MoveRightScale: true,
	}.Update(pressed)

	runtime.GC()
	return nil
}

var _debugColor = color.RGBA{G: 100, A: 100}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Opponent.Draw(screen)
	g.Player.Draw(screen)
	g.PlayerWeapon.Draw(screen)

	g.pikachuAnimeImg.Draw(screen)

	g.GameInfo.Draw(screen)
	for _, w := range g.Walls {
		w.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
