package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 7, 8, 9}
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}
	println()
	for i, num := range nums {
		fmt.Println(i, num)
	}
	mp := map[string]string{"raju": "dhdj", "shd": "dhdh"}
	for k, val := range mp {
		fmt.Println(k, val)
	}
	for k, val := range "golang" {
		fmt.Println(k, val)
	}
	for k, val := range "golang" {
		fmt.Println(k, string(val))
	}
}
