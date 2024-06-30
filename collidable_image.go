package ebitenpkg

import sysimage "image"

type collidableImage struct {
	Collidable
	Image
}

func NewCollidableImage(s Space, t BodyType, i sysimage.Image, a ...Align) CollidableImage {
	img := NewImage(i, a...)

	b := &collidableImage{
		Collidable: newCollidable(s, t, img.GetController()),
		Image:      img,
	}

	s.AddBody(b)

	return b
}

/*
	embedController
*/

func (ci *collidableImage) Align(a Align) CollidableImage {
	ci.Image.Align(a)
	return ci
}

func (ci *collidableImage) Move(x, y float64, replace ...bool) CollidableImage {
	ci.Image.Move(x, y, replace...)
	return ci
}

func (ci *collidableImage) Rotate(degree float64, replace ...bool) CollidableImage {
	ci.Image.Rotate(degree, replace...)
	return ci
}

func (ci *collidableImage) Scale(x, y float64, replace ...bool) CollidableImage {
	ci.Image.Scale(x, y, replace...)
	return ci
}

func (ci *collidableImage) updateControllerReference() CollidableImage {
	ci.Image = ci.Image.updateControllerReference()
	return ci
}

/*
	CollidableImage
*/

func (ci *collidableImage) GetImage() Image {
	return ci.Image
}
