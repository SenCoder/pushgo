package main

import (
	"fmt"
	"unsafe"
)

// 定义非空结构体
type S struct {
	a uint16
	b uint32
}


// 空结构体
var Exists = struct{}{}
// Set is the main interface
type Set struct {
	// struct为结构体类型的变量
	m map[interface{}]struct{}
}


func test1() {
	var s S
	fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6
	var s2 struct{}
	fmt.Println(unsafe.Sizeof(s2)) // prints 0
}


func test2() {
	a := struct{}{}
	b := struct{}{}
	fmt.Println(a == b) // true
	fmt.Printf("%p, %p\n", &a, &b)
}

func test3() {
	s := Set{m:make(map[interface{}]struct{})}
	s.m[12] = Exists
	s.m[13] = Exists
	s.m["hello"] = Exists

	fmt.Printf("%#v", s)
}

func test4() {
	s := struct {
		Name string
	}{Name:"SenCoder"}

	fmt.Println(s)
}


func main() {
	test1()
	fmt.Println("===============")
	test2()
	fmt.Println("===============")
	test3()
}