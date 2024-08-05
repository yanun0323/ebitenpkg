package ebitenpkg

type Animatable interface {
	Drawable
	UpdateFrame()
	Current() Image
}
