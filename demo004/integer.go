package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var a int32 = 5
	var b rune
	b = a
	fmt.Println(b)

	var c uint8 = 3
	var d byte
	d = c
	fmt.Println(d)

	// 0000 0101
	var e uint8 = 5
	// 0000 1001
	var f uint8 = 9
	// 0000 0101
	// 对比f
	// 0000 0100
	// 即是 4
	fmt.Println(e &^ f)

	g := 1e100
	h := int(g)
	fmt.Println(h)

	i := 'a'
	j := '国'
	k := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", i)
	fmt.Printf("%d %[1]c %[1]q\n", j)
	fmt.Printf("%d %[1]c %[1]q\n", k)

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt8)
	fmt.Println(math.MinInt16)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

	l := 2 + 3i
	fmt.Println(reflect.TypeOf(l))
	var m complex64 = 4 + 5i
	fmt.Println(reflect.TypeOf(m))
}
