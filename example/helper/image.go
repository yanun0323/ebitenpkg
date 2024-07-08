package helper

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func GopherImage() image.Image {
	f, err := os.Open("./example/helper/gopher.png")
	if err != nil {
		panic(fmt.Sprintf("read file go.png, err: %+v", err))
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(fmt.Sprintf("decode png, err: %+v", err))
	}

	return img
}
