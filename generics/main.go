package main

import (
	"fmt"
	"math"
)

func main() {
	// With generics, you can declare and use functions or types that
	// are written to work with any of a set of types provided by calling code.

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

	/// testing my implementation
	fmt.Printf("Generic square areas: %v and %v \n",
		calculateSquareArea[int, int](4),
		calculateSquareArea[float64, float64](4.54))

	fmt.Printf("Generic circle areas: %v and %v \n",
		calculateCircleArea[int](7),
		calculateCircleArea[float64](7.00))

	fmt.Printf("Generic Rectangle areas: %v and %v \n",
		calculateRectangleArea(7, 3),
		calculateRectangleArea(7.20, 3))

}

//  the following two functions are non generic functions that each add together the values of a map and return the total.

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Generic function:
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// my generic implementation function calculate square area
func calculateSquareArea[W int | float64, A int | float64](width W) A {
	return A(width * width)
}

// my generic implementation function calculate circle area
func calculateCircleArea[Radius int | float64, Area float64](radius Radius) Area {
	return Area(math.Pi) * Area(radius*radius)
}

// declaring type constraints
type Number interface {
	int | float64 | float32
}

// utilizing type constraints
func calculateRectangleArea[Length Number, Width Number, Area float64](length Length, width Width) Area {
	return Area(length) * Area(width)
}
