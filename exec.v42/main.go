package main

import (
	"runtime"
	"time"
	"fmt"
)

//go:noinline
func add(a, b int) int {
	return a + b
}

func deadloop1() {
	for {
		add(3, 5)
	}
}

func main1() {
	runtime.GOMAXPROCS(1)
	go deadloop1()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}


func dummy() {
	add(3, 5)
}

func deadloop2() {
	for {
		dummy()
	}
}


func main2() {
	runtime.GOMAXPROCS(1)
	go deadloop2()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}

func main() {
	main2()
}