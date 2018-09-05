package main

import (
	"fmt"
)

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	fmt.Println(x == nil)

	Foo(x)

	var y interface{}
	fmt.Println(y == nil)
	Foo(y)
}
