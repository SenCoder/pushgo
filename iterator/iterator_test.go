/*
 * Copyright (c) 2018 PingAn. All rights reserved.
 *
 * Created by yuansheng on 12/15/18 10:58 PM.
 */

package util

import (
	"testing"
	"fmt"
)

type God struct {
	Name string
}


func TestSelector_Iter(t *testing.T) {
	s := ChanIterator{data:[]interface{}{"A", "B", 1}}

	for v := range s.Iter() {
		fmt.Println(v)
	}
}

func TestSelector_Iter2(t *testing.T) {

	d := []God{
		{Name: "Yadang"},
		{Name: "Xia wa"},
	}

	s := NewChanIterator(d)

	for v := range s.Iter() {
		if v == nil {
			break
		}
		fmt.Printf("%v %T\n", v, v)
	}

}
