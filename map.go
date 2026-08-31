package main

import (
	"fmt"
	"maps"
)

func main() {
	mp := make(map[string]int)
	mp["Depto"] = 76
	mp["Soumyo"] = 14
	mp["Aritro"] = 108
	println(mp["Soumyo"])
	println("map_length=", len(mp))
	fmt.Println(mp)
	delete(mp, "Soumyo")
	fmt.Println(mp)
	mp1 := map[string]int{"Irfan": 43, "pritom": 111}
	mp2 := map[string]int{"Irfan": 43, "pritom": 111}
	fmt.Println(mp1)
	fmt.Println(maps.Equal(mp1, mp2))

}
