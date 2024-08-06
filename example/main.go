package main

import (
	"fmt"
	"image/color"

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
	LastGC       uint64
	Count        int
	Space        ebitenpkg.Space
	Walls        []ebitenpkg.Image
	PlayerSprite ebitenpkg.Image
	Player       ebitenpkg.Image
	PlayerWeapon ebitenpkg.Image
	Opponent     ebitenpkg.Image
	GameInfo     ebitenpkg.Text

	PikachuAnime ebitenpkg.Image
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

	playerSprite := ebitenpkg.
		NewImage(helper.PikachuSpriteImage()).
		Align(ebitenpkg.AlignCenter).
		Move(400, 400).
		Scale(5, 5).
		Collidable(space, TypePlayer).
		Spriteable(ebitenpkg.SpriteSheetOption{
			SpriteWidthCount:  1,
			SpriteHeightCount: 6,
			SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX, indexY, scaleX, scaleY int) {
				x, y := 0, (timestamp/5)%2
				sX, sY := 1, 1

				switch {
				case direction&ebitenpkg.DirectionUp != 0:
					y += 2
				case direction&ebitenpkg.DirectionDown != 0:
					y += 4
				case direction&ebitenpkg.DirectionRight != 0:
					sX = -1
				}

				return x, y, sX, sY
			},
		})

	playerName := ebitenpkg.NewText("Pikachu", 20).
		Align(ebitenpkg.AlignTop).
		Move(50, 0).
		SetColor(color.White).
		Attach(playerSprite)

	_ = playerName

	weapon := ebitenpkg.NewImage(ebiten.NewImage(100, 30)).
		Align(ebitenpkg.AlignCenter).
		Move(20, 10).Rotate(-30).Attach(player).
		Collidable(space, TypePlayer).
		Debug(true)

	pikachuAnime := ebitenpkg.NewImage(helper.PikachuAnimeImage()).
		Align(ebitenpkg.AlignTopLeading).
		Moving(100, 100, 600).
		Scaling(5, 5, 600).
		Rotating(60, 300).
		Spriteable(ebitenpkg.SpriteSheetOption{
			SpriteWidthCount:  1,
			SpriteHeightCount: 6,
			SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX, indexY, scaleX, scaleY int) {
				return 0, (timestamp / 5) % 6, 1, 1
			},
		}).
		Debug()

	return &Game{
		Space:        space,
		PlayerSprite: playerSprite,
		Player:       player,
		PlayerWeapon: weapon,
		Opponent:     ebitenpkg.NewImage(helper.GopherImage()).Align(ebitenpkg.AlignTop).Move(200, 200).Scale(-1, 1).Collidable(space, TypeOpponent).Debug(true),
		Walls: []ebitenpkg.Image{
			ebitenpkg.NewImage(ebiten.NewImage(10, int(fH))).Align(ebitenpkg.AlignTop).Move(20, 0).Collidable(space, TypeWall).Debug(true),
			ebitenpkg.NewImage(ebiten.NewImage(10, int(fH))).Align(ebitenpkg.AlignTop).Move(fW-20, 0).Collidable(space, TypeWall).Debug(true),
		},
		GameInfo:     ebitenpkg.NewText("Hello, World!", 20).Align(ebitenpkg.AlignTopLeading).Move(10, 0).SetColor(color.RGBA{R: 100, G: 100, B: 100, A: 100}).SetLineSpacing(50).Debug(true),
		PikachuAnime: pikachuAnime,
	}
}

func (g *Game) Update() error {
	g.Count++
	ebitenpkg.GameUpdate()
	g.Space.GameUpdate()
	g.GameInfo.SetText(fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	pressed := inpututil.AppendPressedKeys(nil)

	g.Player.Debug(g.Space.IsCollided(g.Player))
	g.PlayerWeapon.Debug(g.Space.IsCollided(g.PlayerWeapon))
	g.Opponent.Debug(g.Space.IsCollided(g.Opponent))

	helper.InputHandler[ebitenpkg.Image]{
		Object:         g.PlayerSprite,
		MoveUp:         true,
		MoveDown:       true,
		MoveLeft:       true,
		MoveLeftScale:  false,
		MoveRight:      true,
		MoveRightScale: false,
		RotateLeft:     true,
		RotateRight:    true,
	}.Update(pressed)

	// helper.InputHandler[ebitenpkg.Image]{
	// 	Object:         g.Player,
	// 	MoveUp:         true,
	// 	MoveDown:       true,
	// 	MoveLeft:       true,
	// 	MoveLeftScale:  true,
	// 	MoveRight:      true,
	// 	MoveRightScale: true,
	// 	RotateLeft:     true,
	// 	RotateRight:    true,
	// }.Update(pressed)

	// helper.InputHandler[ebitenpkg.Image]{
	// 	Object:         g.PlayerWeapon,
	// 	MoveLeftScale:  true,
	// 	MoveRightScale: true,
	// }.Update(pressed)

	// ms := runtime.MemStats{}
	// runtime.ReadMemStats(&ms)

	// ts := time.Now().UnixNano() - int64(g.LastGC)
	// println("alloc", ms.Alloc/1024, "kb", ts/1e6, "ms")
	// if g.LastGC != ms.LastGC {
	// 	g.LastGC = ms.LastGC
	// 	println("GC!!!!")
	// }

	// runtime.GC()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Opponent.Draw(screen)
	g.Player.Draw(screen)
	g.PlayerWeapon.Draw(screen)

	g.PlayerSprite.Draw(screen)
	g.PikachuAnime.Draw(screen)

	g.GameInfo.Draw(screen)
	for _, w := range g.Walls {
		w.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
