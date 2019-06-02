package main

import "fmt"

func main() {
	str := "hello world!"
	s := str[:]
	fmt.Printf("%p, %p\n", &str, &s)

	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
		go command [arguments]
	...`
	fmt.Println(GoUsage)
}
