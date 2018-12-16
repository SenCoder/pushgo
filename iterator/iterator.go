/*
 * Copyright (c) 2018 PingAn. All rights reserved.
 *
 * Created by yuansheng on 12/16/18 10:22 AM.
 */

package util

import (
	"reflect"
	"fmt"
	"sync"
)

type Iterator interface {

	Iter() <-chan interface{}
}



type ChanIterator struct {
	sync.RWMutex
	data []interface{}
}

func NewChanIterator(args interface{}) *ChanIterator {

	val := reflect.ValueOf(args)

	if val.Kind() != reflect.Slice {
		panic(fmt.Sprintf("invalid Kind: %v", val.Kind()))
	}

	var data[]interface{}
	for i := 0; i < val.Len(); i++ {
		data = append(data, val.Index(i).Interface())
	}

	return &ChanIterator{data:data}
}

func (i *ChanIterator) Iter() <-chan interface{} {
	ch := make(chan interface{}, len(i.data))

	go func() {
		i.RLock()

		for _, e := range i.data {
			ch <- e
		}

		close(ch)
		i.RUnlock()
	}()

	return ch
}
