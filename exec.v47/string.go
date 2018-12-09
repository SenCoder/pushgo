package main

import (
	"reflect"
	"unsafe"
)

type Bytes []byte

type String string


// unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象，uintptr类型的目标会被回收

func StringBytes(s string) Bytes {
	var bh reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return *(*Bytes)(unsafe.Pointer(&bh))
}


func BytesString(b []byte) String {
	return *(*String)(unsafe.Pointer(&b))
}
