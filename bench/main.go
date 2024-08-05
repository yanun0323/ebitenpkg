package bench

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yanun0323/ebitenpkg"
)

func main() {
	type Data struct{}

	g := &ebitenpkg.Game[Data]{
		Data: Data{},
		Updates: func(d Data) error {
			return nil
		},
		Draws: func(data Data, screen *ebiten.Image) {
		},
		Layouts: func(data Data, outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
			return 1024, 768
		},
	}

	if err := g.Run(); err != nil {
		panic(err)
	}
}
