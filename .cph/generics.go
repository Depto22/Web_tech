package main

import "fmt"
func print_Int_Slice(items []int){
	for _,item:=range items{
		fmt.Print(item," ")
	}
	println()
}
func print_String_Slice(items []string){
	for _,item:=range items{
		fmt.Print(item," ")
	}
	println()
}
//if we write( T int| string) in place of (T any) then it will recieve only int or string slice
//it will give error if we pass boolean slice
func print_Slice[T any](items []T){
	for _,item:=range items{
		fmt.Print(item," ")
	}
	println()
}
// type stack struct{
// 	elements[]int
// }
type stack[T any]struct{
	elements[]T
}
func main(){
	nums:=[]int{1,2,3}
	names:=[]string{"Depto","Irfan","Soumyo","tusher"}
	print_Int_Slice(nums)
	print_String_Slice(names)
	//pass to the function with generics
	print_Slice(nums)
	print_Slice(names)
	myStack:=stack[int]{
		elements:nums,
	}
	myStack1:=stack[string]{
		elements:names,
	}
	fmt.Println(myStack)
	fmt.Println(myStack1)
}