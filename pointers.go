package main

import "fmt"

func changeNum(n *int) {
	*n = 10
	fmt.Println("val in changeNum func:", *n)
}
func main() {
	num := 5
	changeNum(&num)
	fmt.Println("val in changeNum func:", num)
}
