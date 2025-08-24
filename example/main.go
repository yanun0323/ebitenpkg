package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yanun0323/ebitenpkg"
	"github.com/yanun0323/ebitenpkg/example/helper"
)

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
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

	gopher := ebitenpkg.NewImage(helper.GopherImage()).
		Align(ebitenpkg.AlignTopLeading).
		Spriteable(ebitenpkg.SpriteSheetOptionCounter(1, 1, 1, func(fps, timestamp int, direction ebitenpkg.Direction) (indexX, scaleX, scaleY int) {
			return 0, 1, 1
		})).
		Move(300, 300).
		Moving(50, 100, 3*60, true).
		Color(255, 0, 0, 0).
		Coloring(255, 255, 255, 255, 60).
		Collidable(space, TypeOpponent).
		Scale(1, 1)

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
		Spriteable(ebitenpkg.SpriteSheetOptionCounter(1, 6, 6, func(fps, timestamp int, direction ebitenpkg.Direction) (index, scaleX, scaleY int) {
			idx := (timestamp / 5) % 2
			sX, sY := 1, 1

			switch {
			case direction&ebitenpkg.DirectionUp != 0:
				idx += 2
			case direction&ebitenpkg.DirectionDown != 0:
				idx += 4
			case direction&ebitenpkg.DirectionRight != 0:
				sX = -1
			}

			return idx, sX, sY
		}))

	pikachuIdle := ebitenpkg.NewImage(helper.PikachuAnimeImage()).
		Align(ebitenpkg.AlignCenter).
		Move(200, 200).
		Moving(150, 150, 600).
		Scaling(2, 2, 300).
		Rotating(-20, 300).
		Collidable(space, TypeOthers).
		Spriteable(ebitenpkg.SpriteSheetOptionCounter(24, 3, 58, func(fps, timestamp int, direction ebitenpkg.Direction) (index, scaleX, scaleY int) {
			return (timestamp / 6) % fps, 1, 1
		})).
		WithAnimation(ebitenpkg.AnimationEaseInOut())

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
