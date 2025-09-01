package ebitenpkg

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageGroup struct {
	group []Image
}

func NewImageGroup(images ...Image) *ImageGroup {
	return &ImageGroup{
		group: images,
	}
}

func (ig *ImageGroup) UnGroup() []Image {
	g := ig.group
	ig.group = nil
	return g
}

func (ig *ImageGroup) Insert(img Image) {
	ig.group = append(ig.group, img)
}

func (ig *ImageGroup) Draw(screen *ebiten.Image) {
	for _, image := range ig.group {
		image.Draw(screen)
	}
}

func (ig *ImageGroup) Align(align Align) Image {
	for _, image := range ig.group {
		image.Align(align)
	}
	return ig
}

func (ig *ImageGroup) Move(x, y float64, replace ...bool) Image {
	for _, image := range ig.group {
		image.Move(x, y, replace...)
	}
	return ig
}

func (ig *ImageGroup) Moving(x, y float64, tick int, replace ...bool) Image {
	for _, image := range ig.group {
		image.Moving(x, y, tick, replace...)
	}
	return ig
}

func (ig *ImageGroup) Scale(x, y float64, replace ...bool) Image {
	for _, image := range ig.group {
		image.Scale(x, y, replace...)
	}
	return ig
}

func (ig *ImageGroup) Scaling(x, y float64, tick int, replace ...bool) Image {
	for _, image := range ig.group {
		image.Scaling(x, y, tick, replace...)
	}
	return ig
}

func (ig *ImageGroup) Rotate(angle float64, replace ...bool) Image {
	for _, image := range ig.group {
		image.Rotate(angle, replace...)
	}
	return ig
}

func (ig *ImageGroup) Rotating(angle float64, tick int, replace ...bool) Image {
	for _, image := range ig.group {
		image.Rotating(angle, tick, replace...)
	}
	return ig
}

func (ig *ImageGroup) Color(r, g, b, a uint8) Image {
	for _, image := range ig.group {
		image.Color(r, g, b, a)
	}
	return ig
}

func (ig *ImageGroup) Coloring(r, g, b, a uint8, tick int) Image {
	for _, image := range ig.group {
		image.Coloring(r, g, b, a, tick)
	}
	return ig
}

func (ig *ImageGroup) Mask(x, y, w, h float64) Image {
	for _, image := range ig.group {
		image.Mask(x, y, w, h)
	}
	return ig
}

func (ig *ImageGroup) Masking(x, y, w, h float64, tick int) Image {
	for _, image := range ig.group {
		image.Masking(x, y, w, h, tick)
	}
	return ig
}

func (ig *ImageGroup) Spriteable(opt SpriteSheetOption) Image {
	for _, image := range ig.group {
		image.Spriteable(opt)
	}
	return ig
}

func (ig *ImageGroup) Attach(parent Attachable) Image {
	for _, image := range ig.group {
		image.Attach(parent)
	}
	return ig
}

func (ig *ImageGroup) Detach() {
	for _, image := range ig.group {
		image.Detach()
	}
}

func (ig *ImageGroup) Collidable(space Space, group int) Image {
	for _, image := range ig.group {
		image.Collidable(space, group)
	}
	return ig
}

func (ig *ImageGroup) Debug(on ...bool) Image {
	for _, image := range ig.group {
		image.Debug(on...)
	}
	return ig
}

func (ig *ImageGroup) HandleImage(handler func(*ebiten.Image)) Image {
	for _, image := range ig.group {
		image.HandleImage(handler)
	}
	return ig
}

func (ig *ImageGroup) WithAnimation(animation Animation) Image {
	for _, image := range ig.group {
		image.WithAnimation(animation)
	}
	return ig
}

func (ig *ImageGroup) Animation() Animation {
	for _, image := range ig.group {
		return image.Animation()
	}
	return nil
}

func (ig *ImageGroup) Bounds() (width int, height int) {
	for _, image := range ig.group {
		w, h := image.Bounds()
		width = max(width, w)
		height = max(height, h)
	}
	return width, height
}

func (ig *ImageGroup) Aligned() Align {
	for _, image := range ig.group {
		return image.Aligned()
	}
	return 0
}

func (ig *ImageGroup) Moved() (x, y float64) {
	for _, image := range ig.group {
		return image.Moved()
	}
	return 0, 0
}

func (ig *ImageGroup) Scaled() (x, y float64) {
	for _, image := range ig.group {
		return image.Scaled()
	}
	return 0, 0
}

func (ig *ImageGroup) Rotated() (angle float64) {
	for _, image := range ig.group {
		return image.Rotated()
	}
	return 0
}

func (ig *ImageGroup) Colored() (r, g, b, a uint8) {
	for _, image := range ig.group {
		return image.Colored()
	}
	return 0, 0, 0, 0
}

func (ig *ImageGroup) Masked() (x, y, w, h float64) {
	for _, image := range ig.group {
		return image.Masked()
	}
	return 0, 0, 1, 1
}

func (ig *ImageGroup) Debugged() bool {
	for _, image := range ig.group {
		return image.Debugged()
	}
	return false
}

func (ig *ImageGroup) ID() ID {
	for _, image := range ig.group {
		return image.ID()
	}
	return *new(ID)
}

func (ig *ImageGroup) Group() int {
	for _, image := range ig.group {
		return image.Group()
	}
	return 0
}

func (ig *ImageGroup) Parent() Attachable {
	for _, image := range ig.group {
		return image.Parent()
	}
	return nil
}

func (ig *ImageGroup) IsClick(x, y float64) bool {
	for _, image := range ig.group {
		return image.IsClick(x, y)
	}
	return false
}

func (ig *ImageGroup) Copy() Image {
	result := &ImageGroup{}
	for _, image := range ig.group {
		result.Insert(image.Copy())
	}
	return result
}
