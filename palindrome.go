package main

import (
	"fmt"
	"math"
	"time"
)

func avg(x float64, y float64) float64 {
	return float64(x+y) / 2.0
}
func multiple_return(x int, y string) (int, string) {
	return x, y
}
func main() {
	fmt.Println(time.Now())
	n := 65
	fmt.Println(n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%q\n", n)
	fmt.Printf("%u\n", n)
	fmt.Println(math.Pi)
	fmt.Printf("%g\n", avg(3.99, 4.00))
	a, b := multiple_return(78, "sttd")
	fmt.Println(a)
	fmt.Println(b)

}
