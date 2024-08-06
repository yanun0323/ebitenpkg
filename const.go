package ebitenpkg

import (
	"image/color"
	"math"
	"sync"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

const (
	_radianToDegree = 180 / math.Pi
	_floatFix       = 1e-3
)

type ID uuid.UUID

func newID() ID {
	return ID(uuid.New())
}

// func newValue(v any) *atomic.Value {
// 	value := &atomic.Value{}
// 	value.Store(v)
// 	return value
// }

/*
	Game
*/

var (
	_currentGameTimeLock = &sync.RWMutex{}
	_currentGameTime     = 0
)

func GameUpdate() {
	_currentGameTimeLock.Lock()
	defer _currentGameTimeLock.Unlock()

	_currentGameTime++
}

func CurrentGameTime() int {
	_currentGameTimeLock.RLock()
	defer _currentGameTimeLock.RUnlock()

	return _currentGameTime
}

/*
	Default Align
*/

var _defaultAlignValue = newValue(AlignCenter)

func SetDefaultAlign(align Align) Align {
	_defaultAlignValue.Store(align)
	return DefaultAlign()
}

func DefaultAlign() Align {
	return _defaultAlignValue.Load()
}

/*
	Default Text Dpi
*/

var _defaultTextDpi = newValue(72.0)

func SetDefaultTextDpi(dpi float64) {
	_defaultTextDpi.Store(dpi)
}

func DefaultTextDpi() float64 {
	return _defaultTextDpi.Load()
}

/*
	Default Debug Color
*/

var _defaultDebugColor = newValue(color.RGBA64{G: 0xffff >> 1, A: 0xffff >> 1})

func SetDefaultDebugColor(clr color.Color) {
	r, g, b, a := clr.RGBA()
	_defaultDebugColor.Store(color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)})
}

func DefaultDebugColor() color.Color {
	return _defaultDebugColor.Load()
}

/*
	Default Font
*/

var _defaultFont = newValue(fonts.MPlus1pRegular_ttf)

func SetDefaultFont(font []byte) {
	_defaultFont.Store(font)
}

func DefaultFont() []byte {
	return _defaultFont.Load()
}
