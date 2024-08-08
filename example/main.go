package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yanun0323/ebitenpkg"
	"github.com/yanun0323/ebitenpkg/example/helper"
	"github.com/yanun0323/pkg/logs"
)

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		logs.Fatal(err)
	}
}

type Game struct {
	LastGC        uint64
	Space         ebitenpkg.Space
	Walls         []ebitenpkg.Image
	PikachuSprite ebitenpkg.Image
	PikachuIdle   ebitenpkg.Image
	Gopher        ebitenpkg.Image
	GameInfo      ebitenpkg.Text
}

const (
	TypeWall int = iota
	TypePlayer
	TypeOpponent
	TypeOthers
)

func NewGame() ebiten.Game {
	space := ebitenpkg.NewSpace()
	w, h := ebiten.WindowSize()
	fW, fH := float64(w), float64(h)

	layer1 := ebitenpkg.NewRectangle(10, 10, color.White).Align(ebitenpkg.AlignTopLeading).Move(1, 1)
	layer2 := ebitenpkg.NewRectangle(30, 30, color.RGBA{111, 0, 0, 111},
		layer1,
	).Align(ebitenpkg.AlignTopLeading).Move(3, 3)
	layer3 := ebitenpkg.NewRectangle(50, 50, color.RGBA{111, 0, 111, 111},
		layer2,
	).Align(ebitenpkg.AlignTopLeading).Move(25, 25)

	gopher := ebitenpkg.NewRoundedRectangle(100, 100, 15, color.RGBA{0, 0, 111, 111},
		layer3,
	).
		Align(ebitenpkg.AlignTopLeading).
		Move(300, 300).
		Moving(50, 100, 600, true).
		Collidable(space, TypeOpponent)

	pikachuSprite := ebitenpkg.
		NewImage(helper.PikachuSpriteImage(),
			ebitenpkg.NewRectangle(20, 20, color.RGBA{0, 0, 255, 255},
				ebitenpkg.NewText("Pikachu", 20).
					Align(ebitenpkg.AlignBottom).
					Move(0, -30).
					SetColor(color.White),
			).
				Align(ebitenpkg.AlignBottom).
				Move(0, -30),
			ebitenpkg.NewText("---", 20).
				Align(ebitenpkg.AlignTop).
				Move(0, 10).
				SetColor(color.White),
		).
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

	pikachuIdle := ebitenpkg.NewImage(helper.PikachuAnimeImage()).
		Align(ebitenpkg.AlignCenter).
		Move(200, 200).
		Moving(150, 150, 600).
		Scaling(2, 2, 300).
		Rotating(-20, 300).
		Collidable(space, TypeOthers).
		Spriteable(ebitenpkg.SpriteSheetOption{
			SpriteWidthCount:  1,
			SpriteHeightCount: 6,
			SpriteHandler: func(fps, timestamp int, direction ebitenpkg.Direction) (indexX, indexY, scaleX, scaleY int) {
				return 0, (timestamp / 5) % 6, 1, 1
			},
		})

	return &Game{
		Space:         space,
		Gopher:        gopher,
		PikachuSprite: pikachuSprite,
		PikachuIdle:   pikachuIdle,
		Walls: []ebitenpkg.Image{
			ebitenpkg.NewImage(ebiten.NewImage(10, int(fH))).Align(ebitenpkg.AlignTop).Move(20, 0).Collidable(space, TypeWall),
			ebitenpkg.NewImage(ebiten.NewImage(10, int(fH))).Align(ebitenpkg.AlignTop).Move(fW-20, 0).Collidable(space, TypeWall),
		},
		GameInfo: ebitenpkg.NewText("Hello, World!", 20).Align(ebitenpkg.AlignTopLeading).Move(10, 0).SetColor(color.RGBA{R: 100, G: 100, B: 100, A: 100}).SetLineSpacing(50),
	}
}

func (g *Game) Update() error {
	/* update game ticker */
	ebitenpkg.GameUpdate()
	g.Space.GameUpdate()

	/* update game info */
	g.GameInfo.SetText(fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))

	/* handle keyboard input event*/
	pressed := inpututil.AppendPressedKeys(nil)
	helper.InputHandler[ebitenpkg.Image]{
		Object:         g.PikachuSprite,
		MoveUp:         true,
		MoveDown:       true,
		MoveLeft:       true,
		MoveLeftScale:  false,
		MoveRight:      true,
		MoveRightScale: false,
		RotateLeft:     true,
		RotateRight:    true,
	}.Update(pressed)

	/* handle collision debug */
	// g.Gopher.Debug(g.Space.IsCollided(g.Gopher))
	// g.PikachuSprite.Debug(g.Space.IsCollided(g.PikachuSprite))
	// g.PikachuIdle.Debug(g.Space.IsCollided(g.PikachuIdle))
	// for _, w := range g.Walls {
	// 	w.Debug(g.Space.IsCollided(w))
	// }

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/* draw game objects */
	logs.Debug("draw")
	g.Gopher.Draw(screen)
	g.PikachuSprite.Draw(screen)

	g.PikachuIdle.Draw(screen)
	g.GameInfo.Draw(screen)
	for _, w := range g.Walls {
		w.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
