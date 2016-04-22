package sorttester

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(l []int) []int {
	if llen := len(l); llen > 1 {
		l1 := mergeSort(l[:llen/2])
		l2 := mergeSort(l[llen/2:])
		r := make([]int, 0)

		for {
			if len(l1) > 0 && len(l2) > 0 {
				if l1[0] < l2[0] {
					r = append(r, l1[0])
					l1 = l1[1:]
				} else {
					r = append(r, l2[0])
					l2 = l2[1:]
				}
			}

			if len(l1) == 0 {
				r = append(r, l2...)
				break
			}

			if len(l2) == 0 {
				r = append(r, l1...)
				break
			}
		}

		return r
	} else {
		return l
	}
}

func TestMergeSort() {
	N := 100
	l := make([]int, N)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := range l {
		l[i] = rand.Intn(1000)
	}

	fmt.Println(l)

	sl := mergeSort(l)

	fmt.Println(sl)
}
