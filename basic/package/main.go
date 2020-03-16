package main

// import "fmt"
// import "math"
//In Go, a name is exported if it begins with a capital letter.
//For example, Pizza is an exported name, as is Pi, which is exported from the math package.
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(100))
}
