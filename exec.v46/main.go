package main

import "fmt"

// 考点1: golang 中 map 的遍历是随机的
// 考点2: golang 中 range 循环取出的是 map 键值对的副本

func main() {

	m1 := make(map[int]int)
	m2 := make(map[int]*int)

	m1[1] = 1
	m1[2] = 2
	m1[3] = 3
	m1[4] = 4
	m1[5] = 5

	for k, v := range m1 {

		m2[k] = &v
		fmt.Println(v, "======")
	}

	fmt.Println(m2)
	for _, v := range m2 {
		fmt.Println(*v)
	}

}
