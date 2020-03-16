package main

import (
	"fmt"
	//"math"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	b := [...]int{1, 2, 4, 5, 6, 7}
	fmt.Println(b)
	fmt.Printf("a:%T\n", a)
	fmt.Printf("primes:%T\n", primes)
	fmt.Printf("b:%T\n", b)
}
