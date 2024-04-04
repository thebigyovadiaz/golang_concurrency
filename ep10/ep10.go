package ep10

import (
	"fmt"
	"math"
)

func EP10() {
	arr := []int{-15, -7, -9, -14, -12}
	numb := computeClosestToZero(arr)
	fmt.Println("number: ", numb)
}

func computeClosestToZero(ts []int) int {
	nearPositive := 1
	nearNegative := -1
	zero := 0

	fmt.Println(ts)

	for _, v := range ts {
		if v < zero && v <= nearPositive {
			nearNegative = v
		} else if v >= zero && v <= nearPositive {
			nearPositive = v
		}
	}

	absNeg := int(math.Abs(float64(nearNegative)))

	if absNeg == nearPositive {
		return nearPositive
	}

	return -1
}
