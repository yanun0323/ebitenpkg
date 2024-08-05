package ebitenpkg

import (
	"errors"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestGame(t *testing.T) {
	type TestData struct {
		Shutdown chan bool
	}

	testErr := errors.New("stop game for test")

	g := Game[*TestData]{
		Data: &TestData{
			Shutdown: make(chan bool, 1),
		},
		Updates: func(data *TestData) error {
			select {
			case <-data.Shutdown:
				t.Log("shutdown signal received")
				return testErr
			default:
				return nil
			}
		},
		Draws: func(data *TestData, screen *ebiten.Image) {
		},
		Layouts: func(data *TestData, outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
			return 100, 100
		},
	}

	_ = g

	// go func() {
	// 	<-time.After(1 * time.Second)
	// 	t.Log("send shutdown signal")
	// 	g.Data.Shutdown <- true
	// }()

	// if err := g.Run(); err != nil {
	// 	if errors.Is(err, testErr) {
	// 		t.Log("test passed")
	// 	} else {
	// 		t.Fatal(err)
	// 	}
	// }
}
