package ep10

import (
	"fmt"
	"math"
)

func EP10() {
	arr := []int{-1, -7, -1, 5, 1}
	numb := computeClosestToZero(arr)
	fmt.Println("number: ", numb)
}

// computeClosestToZero find a number near zero
func computeClosestToZero(ts []int) int {
	var numb, key int

	fmt.Println(ts)

	for k, v := range ts {
		nABS := int(math.Abs(float64(v)))

		if k == 0 {
			numb = nABS
			key = k
		} else {
			if numb > nABS {
				numb = nABS
				key = k
			} else if numb == nABS && ts[key] < v {
				key = k
			} else {
				numb = nABS
				key = k
			}
		}
	}

	if len(ts) == 0 {
		return 0
	}

	fmt.Println("numb:", numb, key)
	return ts[key]
}
