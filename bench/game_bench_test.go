package bench

import (
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type TestGame struct {
	Img      *ebiten.Image
	Shutdown chan struct{}
}

func NewGame() *TestGame {
	ch := make(chan struct{})
	go func() {
		<-time.After(time.Second)
		ch <- struct{}{}
	}()

	return &TestGame{
		Img:      ebiten.NewImage(500, 500),
		Shutdown: ch,
	}
}

func (g *TestGame) Update() error {
	println("update")
	select {
	case <-g.Shutdown:
		return ebiten.Termination
	default:
		return nil
	}
}

func (g *TestGame) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Img, nil)
}

func (g *TestGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func BenchmarkEbitenGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := NewGame()
		go ebiten.RunGame(g)
	}
}

func BenchmarkEbitenpkgGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := ebitenpkg.Game[*TestGame]{
			Data: NewGame(),
			Updates: func(d *TestGame) error {
				println("update")
				select {
				case <-d.Shutdown:
					return ebiten.Termination
				default:
					return nil
				}
			},
			Draws: func(data *TestGame, screen *ebiten.Image) {
				screen.DrawImage(screen, nil)
			},
			Layouts: func(data *TestGame, outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
				return 1024, 768
			},
		}

		go g.Run()
	}
}
