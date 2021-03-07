package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [3]string{"pos1", "pos2", "pos3"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3, reflect.TypeOf(array3))
	fmt.Println("")
	fmt.Println("Slices")
	slice := []int{}
	fmt.Println(slice, reflect.TypeOf(slice))
	slice = append(slice, 1)
	fmt.Println(slice)

	slice2 := array2[0:2]
	fmt.Println(slice2)
}
