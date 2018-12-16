/*
 * Copyright (c) 2018 PingAn. All rights reserved.
 *
 * Created by yuansheng on 12/15/18 10:52 PM.
 */

package util

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomIterator struct {
	ChanIterator
}

func NewRandomIterator(args interface{}) *RandomIterator {
	itor := &RandomIterator{ChanIterator: *NewChanIterator(args)}

	itor.Shuffle()
	return itor
}

func (i *RandomIterator) Shuffle() {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for n := len(i.data); n >0; n -- {
		randIndex := r.Intn(n)
		i.data[n-1], i.data[randIndex] = i.data[randIndex] , i.data[n-1]
	}
	fmt.Println(i.data)
}
