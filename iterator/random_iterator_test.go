/*
 * Copyright (c) 2018 PingAn. All rights reserved.
 *
 * Created by yuansheng on 12/16/18 10:39 AM.
 */

package util

import (
	"testing"
)

func TestRandomIterator_Iter(t *testing.T) {

	d := []God{
		{Name: "A"},
	}
	i := NewRandomIterator(d)

	for v := range i.Iter() {
		if v == nil {
			break
		}
	}

	d = []God{
		{Name: "A"},
		{Name: "B"},
	}
	i = NewRandomIterator(d)

	for v := range i.Iter() {
		if v == nil {
			break
		}
	}
}