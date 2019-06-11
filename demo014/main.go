package main

import "fmt"

type Human struct {
	name string
	sex int
}

type Student struct {
	Human
	schoolName string
}

type Helloer interface {
	sayHello()
}

func (h *Human) sayHello() {
	fmt.Println("say Hello")
}

func main() {
	var helloer Helloer
	student := Student{Human{"dimdark", 0}, "scau"}
	// helloer = student 错误
	helloer = &student
	// student.sayHello()
	helloer.sayHello()
}




