package main

import (
	"fmt"
	//"math"
)

func main() {
	i, j := 41, 7001

	p := &i
	fmt.Println(p)  // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(p)
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
