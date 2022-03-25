package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

}

//SumInts adds together the values of sliceInt.
func SumInts(sliceInt map[string]int64) int64 {
	var totalSum int64
	for _, value := range sliceInt {
		totalSum += value
	}
	return totalSum
}

//SumFloadts adds together the values of sliceFloat.
func SumFloats(sliceFloat map[string]float64) float64 {
	var totalSum float64
	for _, value := range sliceFloat {
		totalSum += value
	}
	return totalSum
}

// SumIntsOrFloats sums the values of map s. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[Key comparable, TypeAccepted int64 | float64](myMap map[Key]TypeAccepted) TypeAccepted {
	var sumTotal TypeAccepted
	for _, value := range myMap {
		sumTotal += value
	}
	return sumTotal
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
