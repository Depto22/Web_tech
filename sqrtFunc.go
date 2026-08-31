package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Println(sqrt(n))
}
