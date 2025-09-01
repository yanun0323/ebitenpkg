package ebitenpkg

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type animatableValue[T any] interface {
	Add(T) T
	Sub(T) T
	Ratio(float64) T
	Limit() T
}

/*
	##     ##       ###        ######     ##    ##    ########    ########
	###   ###      ## ##      ##    ##    ##   ##     ##          ##     ##
	#### ####     ##   ##     ##          ##  ##      ##          ##     ##
	## ### ##    ##     ##     ######     #####       ######      ########
	##     ##    #########          ##    ##  ##      ##          ##   ##
	##     ##    ##     ##    ##    ##    ##   ##     ##          ##    ##
	##     ##    ##     ##     ######     ##    ##    ########    ##     ##
*/

// masker is a struct for apply to mask the subImage
type masker struct {
	X float64 // range[0,1]
	Y float64 // range[0,1]
	W float64 // range[0,1]
	H float64 // range[0,1]
}

func (m masker) Apply(img *ebiten.Image) *ebiten.Image {
	if m.NoMask() {
		return img
	}

	b := img.Bounds()
	x := int(m.X * float64(b.Dx()))
	y := int(m.Y * float64(b.Dy()))
	w := int(m.W * float64(b.Dx()))
	h := int(m.H * float64(b.Dy()))
	w = max(w, 0)
	h = max(h, 0)
	return img.SubImage(image.Rect(x, y, x+w, y+h)).(*ebiten.Image)
}

func (m masker) NoMask() bool {
	return m.X <= 0 && m.Y <= 0 && m.W >= 1 && m.H >= 1
}

func (m masker) Empty() bool {
	return m.W <= 0 || m.H <= 0
}

func (m masker) Sub(s masker) masker {
	return masker{
		X: m.X - s.X,
		Y: m.Y - s.Y,
		W: m.W - s.W,
		H: m.H - s.H,
	}
}

func (m masker) Add(s masker) masker {
	return masker{
		X: m.X + s.X,
		Y: m.Y + s.Y,
		W: m.W + s.W,
		H: m.H + s.H,
	}
}

func (m masker) Ratio(f float64) masker {
	return masker{
		X: m.X * f,
		Y: m.Y * f,
		W: m.W * f,
		H: m.H * f,
	}
}

func (m masker) Limit() masker {
	if m.W <= 0 || m.H <= 0 {
		return masker{0, 0, 1, 1}
	}

	return m
}

/*
	########      #######     ########       ###       ########    ####     #######     ##    ##
	##     ##    ##     ##       ##         ## ##         ##        ##     ##     ##    ###   ##
	##     ##    ##     ##       ##        ##   ##        ##        ##     ##     ##    ####  ##
	########     ##     ##       ##       ##     ##       ##        ##     ##     ##    ## ## ##
	##   ##      ##     ##       ##       #########       ##        ##     ##     ##    ##  ####
	##    ##     ##     ##       ##       ##     ##       ##        ##     ##     ##    ##   ###
	##     ##     #######        ##       ##     ##       ##       ####     #######     ##    ##
*/

type rotation float64

func (r rotation) Add(rr rotation) rotation {
	return r + rr
}

func (r rotation) Sub(rr rotation) rotation {
	return r - rr
}

func (r rotation) Ratio(f float64) rotation {
	return rotation(float64(r) * f)
}

func (r rotation) Float64() float64 {
	return float64(r)
}

func (r rotation) Limit() rotation {
	return r
}

/*
	######      #######     ##           #######     ########      ######
	##    ##    ##     ##    ##          ##     ##    ##     ##    ##    ##
	##          ##     ##    ##          ##     ##    ##     ##    ##
	##          ##     ##    ##          ##     ##    ########      ######
	##          ##     ##    ##          ##     ##    ##   ##            ##
	##    ##    ##     ##    ##          ##     ##    ##    ##     ##    ##
	######      #######     ########     #######     ##     ##     ######
*/

type colors [4]uint8

func (c colors) Add(c2 colors) colors {
	return colors{
		uint8(min(max(int(c[0])+int(c2[0]), 0), 255)),
		uint8(min(max(int(c[1])+int(c2[1]), 0), 255)),
		uint8(min(max(int(c[2])+int(c2[2]), 0), 255)),
		uint8(min(max(int(c[3])+int(c2[3]), 0), 255)),
	}
}

func (c colors) Sub(c2 colors) colors {
	return colors{
		uint8(min(max(int(c[0])-int(c2[0]), 0), 255)),
		uint8(min(max(int(c[1])-int(c2[1]), 0), 255)),
		uint8(min(max(int(c[2])-int(c2[2]), 0), 255)),
		uint8(min(max(int(c[3])-int(c2[3]), 0), 255)),
	}
}

func (c colors) Ratio(r float64) colors {
	return colors{
		uint8(min(max(float64(c[0])*r, 0), 255)),
		uint8(min(max(float64(c[1])*r, 0), 255)),
		uint8(min(max(float64(c[2])*r, 0), 255)),
		uint8(min(max(float64(c[3])*r, 0), 255)),
	}
}

func (c colors) Limit() colors {
	c[0] = max(min(c[0], 255), 0)
	c[1] = max(min(c[1], 255), 0)
	c[2] = max(min(c[2], 255), 0)
	c[3] = max(min(c[3], 255), 0)

	return c
}
