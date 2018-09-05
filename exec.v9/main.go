package main

import (
	"fmt"
	"sync"
	"time"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() chan interface{} {
	// ch := make(chan interface{})
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem := range set.s {
			fmt.Println("s ")
			ch <- elem
			fmt.Println("e ")
		}

		close(ch)
		set.RUnlock()

	}()
	fmt.Println("return ")
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2", "3", "4"},
	}
	v := <-th.Iter()
	fmt.Printf("%s %v \n", "ch", v)

	time.Sleep(time.Second)
}
