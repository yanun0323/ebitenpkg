package ebitenpkg

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
