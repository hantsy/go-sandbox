package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())

	v3 := Vertex{3, 4}
	Scale2(&v3, 10)
	fmt.Println(Abs2(v3))

	fmt.Println("with pointers.")
	v4 := Vertex{3, 4}
	v4.Scale(2)
	Scale2(&v4, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	Scale2(p, 8)

	fmt.Println(v4, p)
}
