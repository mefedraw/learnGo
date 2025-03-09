package other

import (
	"fmt"
	"math"
)

func Task1() {
	x := 2.0
	y := 3.0

	result := math.Pow(x, y)

	fmt.Println("%f ^ %f = %f\n", x, y, result) // Printf()
}
