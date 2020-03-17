package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//If i does not hold a T, the statement will trigger a panic.
	f = i.(float64) // panic
	fmt.Println(f)
}
