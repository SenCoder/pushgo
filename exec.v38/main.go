package main

import "fmt"

type Student interface {

	Set(name string)

	Get() string
}


type S struct {
	Name string
}

func (s *S) Set(name string) {
	s.Name = name
}


func (s S) Get() string {
	return s.Name
}

func (s *S) Reset(name string) *S {
	s.Name = name
	return s
}

func (s *S) Play() string {
	fmt.Println(s.Name)
	return s.Name
}


func main() {
	ss := S{}
	ss.Set("yuansheng")

	var stu Student
	stu = &ss

	ss.Reset("peter")
	//S{}.Play()	编译无法通过
	fmt.Println(stu.Get())

	// 编译无法通过
	//fmt.Println(S{"yuansheng"}.Reset("michael").Play())
}