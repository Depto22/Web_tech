package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLang() (string, string) {
	return "nnd", "dddl"
}
func sum(num ...int) int {
	var sum = 0
	for _, val := range num {
		sum += val
	}
	return sum
}
func prnt(num ...interface{}) {

	for k, val := range num {
		fmt.Print(val, " ", k, " ")
	}
	println()

}
func main() {
	nums := []int{3, 7, 8, 9, 22}
	fmt.Println(add(4, 9))
	fmt.Println(getLang())
	fmt.Println(sum(5, 7, 7, 43, 2, 4, 4))
	fmt.Println(sum(5, 7, 7, 43, 2, 4, 4, 3, 1, 4))
	prnt(3, 6.29, "sjhd", 020, 37)
	fmt.Println(sum(nums...))
	//_:=6

}
