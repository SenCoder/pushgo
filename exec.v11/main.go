package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Man interface {
	People
	War()
}

type Student struct{}

type Teacher struct{}

type Soldier struct {
}

func (sd *Soldier) Show() {

}

func (sd *Soldier) War() {

}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func test() {
	sd := &Soldier{}

	a := People(sd)
	c := Man(sd).(*Soldier)
	fmt.Println("%T %T", a, c)

	//b, ok := People(sd).(*Man)
	//if ok {
	//	fmt.Println("%T", b, ok)
	//}


	//stu := &Student{}


}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	var tea *Teacher
	if tea == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	test()
}
