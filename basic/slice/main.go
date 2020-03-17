package main

import (
	"fmt"
	"strings"
)

func changes() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	// new instance
	names1 := names
	// change the value does not affect the original array
	names1[0] = "Change it"

	fmt.Println(names)
	fmt.Println(names1)

	// a slice is a reference of the underlied array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func literals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func defaultBounds() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array,
// counting from the first element in the slice.
func lengths() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func zeros() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeit() {
	a := make([]int, 5) // len(a)=5
	printSlice2("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice2("b", b)

	c := b[:cap(b)] // len(b)=5, cap(b)=5
	printSlice2("c", c)

	d := c[1:] // len(b)=4, cap(b)=4
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func ticTacToe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// The resulting value of append is a slice containing
// all the elements of the original slice plus the provided values.
// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
// The returned slice will point to the newly allocated array.
func appends() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4, 5)
	printSlice(s)

	s = append(s, []int{7, 8}...)
	printSlice(s)
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	changes()
	literals()
	defaultBounds()
	lengths()
	zeros()
	makeit()
	ticTacToe()
	appends()
}
