package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	i := 42
	f := math.Sqrt(float64(i))
	u := uint(f)
	fmt.Println(i, f, u)

	k := 42           // int
	j := k            // j is an int
	pi := 3.142       // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("type of k is: %T \n", k)
	fmt.Printf("type of j is: %T \n", j)
	fmt.Printf("type of pi is: %T \n", pi)
	fmt.Printf("type of g is: %T \n", g)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
