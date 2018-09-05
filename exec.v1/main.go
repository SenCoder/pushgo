package main

import (
	"fmt"
)

// defer/panic 的执行顺序

func main() {
	defer func() { fmt.Println("main defer") }()

	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
