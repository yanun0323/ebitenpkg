package bench

import (
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

type TestEbitenImage struct {
	Img      *ebiten.Image
	Opt      *ebiten.DrawImageOptions
	Shutdown chan struct{}
}

func NewTestEbitenImage() TestEbitenImage {
	ch := make(chan struct{})
	go func() {
		<-time.After(time.Second)
		ch <- struct{}{}
	}()

	return TestEbitenImage{
		Img:      ebiten.NewImage(500, 500),
		Opt:      &ebiten.DrawImageOptions{},
		Shutdown: ch,
	}
}

type TestEbitenpkgImage struct {
	Img      ebitenpkg.Image
	Shutdown chan struct{}
}

func NewTestEbitenpkgImage() TestEbitenpkgImage {
	ch := make(chan struct{})
	go func() {
		<-time.After(time.Second)
		ch <- struct{}{}
	}()

	return TestEbitenpkgImage{
		Img:      ebitenpkg.NewImage(ebiten.NewImage(500, 500)),
		Shutdown: ch,
	}
}

func BenchmarkEbitenImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := ebitenpkg.Game[TestEbitenImage]{
			Data: NewTestEbitenImage(),
			Updates: func(data TestEbitenImage) error {
				select {
				case <-data.Shutdown:
					return ebiten.Termination
				default:
					data.Opt.GeoM.Translate(0, 100)
					data.Opt.GeoM.Scale(2, 2)
					return nil
				}
			},
			Draws: func(data TestEbitenImage, screen *ebiten.Image) {
				screen.DrawImage(data.Img, data.Opt)
			},
			Layouts: func(data TestEbitenImage, outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
				return 1024, 768
			},
		}

		go g.Run()
	}
}

func BenchmarkEbitenpkgImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := ebitenpkg.Game[TestEbitenpkgImage]{
			Data: NewTestEbitenpkgImage(),
			Updates: func(data TestEbitenpkgImage) error {
				select {
				case <-data.Shutdown:
					return ebiten.Termination
				default:
					data.Img.Move(0, 100).Scale(2, 2)
					return nil
				}
			},
			Draws: func(data TestEbitenpkgImage, screen *ebiten.Image) {
				data.Img.Draw(screen)
			},
			Layouts: func(data TestEbitenpkgImage, outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
				return 1024, 768
			},
		}

		go g.Run()
	}
}
