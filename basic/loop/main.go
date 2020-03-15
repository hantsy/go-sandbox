package main

import (
	"fmt"
)

//https://gist.github.com/mannharleen/8f6717f892912d603699b51da4cad2ec
var z = 100.0      // the seed
var delta = 0.0001 // the epsilon

func Sqrt(x float64) float64 {
	z = z - (z*z-x)/(2*z)
	if z*z-x < delta {
		return z
	}
	fmt.Println("value of z = %f", z)
	return Sqrt(x)
}

func main() {
	fmt.Println("Sqrt(2):", Sqrt(2))
	fmt.Println("Sqrt(4):", Sqrt(4))
	fmt.Println("Sqrt(6):", Sqrt(6))
	fmt.Println("Sqrt(1024):", Sqrt(1024))
}
