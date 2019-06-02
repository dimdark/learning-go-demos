package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println("-----------------------------")

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	fmt.Println(string(65))
	fmt.Println(string(0x4eac))
	fmt.Println(string(1234567))
	fmt.Println(string(9999999))

}
