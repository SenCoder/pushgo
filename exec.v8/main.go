package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {

	ua.RLock()
	defer ua.RUnlock()
	fmt.Println("<<")
	defer fmt.Println(">>")

	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := UserAges{ages: make(map[string]int)}

	go func() {
		for i := 0; i < 10; i++ {
			ua.Add("wen", 21)
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(ua.Get("wen"))
		}()
	}

	time.Sleep(time.Second)
}
