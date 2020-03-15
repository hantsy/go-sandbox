package main

import (
	"fmt"
)

func sum() {
	// The basic for loop has three components separated by semicolons:

	// 	the init statement: executed before the first iteration
	// 	the condition expression: evaluated before every iteration
	// 	the post statement: executed at the end of every iteration

	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}

func sum2() {
	//The init and post statements are optional.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sum3() {
	// Go's while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func buzzfoo() {
	for i := 0; i < 100; i++ {

		if i%15 == 0 {
			fmt.Println("buzzfoo")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("foo")
		} else {
			fmt.Println(i)
		}
	}

}

func main() {

	sum()
	sum2()
	sum3()
	buzzfoo()
}
