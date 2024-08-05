package ebitenpkg

type Direction int

const (
	Left  Direction = 1 << 0
	Right Direction = 1 << 1
	Up    Direction = 1 << 2
	Down  Direction = 1 << 3
)
