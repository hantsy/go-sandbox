package main

import "fmt"

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

func main() {
	fmt.Printf(" %d + %d = %d \n", 3, 5, add(3, 5))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(17))
}
