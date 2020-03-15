package main

import "fmt"

var rust, ruby bool

var i, j int = 1, 2

const Pi = 3.14159

func main() {
	fmt.Println(rust, ruby)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	var m int
	//nside a function, the := short assignment statement can be used in place of a var declaration with implicit type. 
	n:= 10
	fmt.Println(m, n)

	
    // 0 for numeric types,
    // false for the boolean type, and
    // "" (the empty string) for strings.

	var j int
	var f float32
	var d float64
	var b bool
	var s string
	fmt.Println(j, f, d, b, s)

	const Truth= true

	fmt.Println(Pi, Truth)

}