package pic

import (
	"golang.org/x/tour/pic"
)

func Test() {
	Pic := func(dx, dy int) [][]uint8 {
		p := make([][]uint8, dy)
		for i := 0; i < dy; i++ {
			p[i] = make([]uint8, dx)
			for j := 0; j < dx; j++ {
				p[i][j] = uint8(i + j)
			}
		}

		return p
	}

	pic.Show(Pic)
}
