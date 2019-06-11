package main

import (
	"fmt"
	"math"
)

const (
	a int = 3
	b int = 4
)

const (
	c = 3
	d = 4
)

func main() {
	// fmt.Println(math.Sqrt(a * a + b * b))
	fmt.Println(math.Sqrt(c * c + d * d))
	fmt.Printf("%T %T %T %T", a, b, c, d)
}





