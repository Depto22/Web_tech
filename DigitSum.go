package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	//sum := 0
	var sum int = 0
	for i := n; i != 0; i /= 10 {
		sum += i % 10
	}
	fmt.Println(sum)
	//var isAdult = true

}
