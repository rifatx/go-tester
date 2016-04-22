package spiral

import (
	"fmt"
	"math"
)

func Draw(r int) {
	_r, fi := 0.0, 0.0

	m := make([][]string, 2*r)

	for i := 0; i < 2*r; i++ {
		m[i] = make([]string, 2*r)

		for j := 0; j < 2*r; j++ {
			m[i][j] = " "
		}
	}

	for _r < float64(r) {
		x := math.Trunc(_r*math.Cos(fi) + float64(r))
		y := math.Trunc(_r*math.Sin(fi) + float64(r))
		m[int(y)][int(x)] = "*"
		_r, fi = _r+0.2, fi+0.2
	}

	for _, s := range m {
		fmt.Println(s)
	}
}
