package helper

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func GopherImage() image.Image {
	return img("gopher")
}

func PikachuAnimeImage() image.Image {
	return img("pikachu_anime")
}

func PikachuSpriteImage() image.Image {
	return img("pikachu_sprite")
}

func img(name string) image.Image {
	f, err := os.Open("./example/helper/" + name + ".png")
	if err != nil {
		panic(fmt.Sprintf("read file %s.png, err: %+v", name, err))
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(fmt.Sprintf("decode png, err: %+v", err))
	}

	return img
}
