package main

import (
	"fmt"
	"math"
)

func main() {
	const intervalInt = ^uint(0) >> 1
	fmt.Println(intervalInt)
	const maxInt = int(^uint(0) >> 1)
	s := fmt.Sprintf("%b", maxInt)
	fmt.Println(len(s))
	fmt.Printf("%b\n", maxInt)
	fmt.Println(maxInt)

	var num int
	num = int(math.MaxInt64)
	fmt.Println(num)
}










