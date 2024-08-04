package ebitenpkg

import sysimage "image"

type SpiritSheetOption func(moveX, moveY float64, imgWidth, imgHeight int, side, currentSecond int) sysimage.Rectangle
