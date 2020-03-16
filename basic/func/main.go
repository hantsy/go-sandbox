package main

import (
	"fmt"
	"math"
)

// A function can take zero or more arguments.
// Notice that the type comes after the variable name.
// https://blog.golang.org/gos-declaration-syntax

// func add(x int, y int ) int {
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

// return multi results.
func swap(x, y string) (string, string) {
	return y, x
}

//named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, next := 0, 1
	return func() int {
		sum := prev + next
		prev, next = next, sum
		return sum
	}
}

func main() {
	fmt.Printf(" %d + %d = %d \n", 3, 5, add(3, 5))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(17))

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// closure
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
