package main

import (
	"testing"
	"fmt"
)

const N = 3000000

// string常量会在编译期分配到只读段，对应数据地址不可写入，并且相同的string常量不会重复存储。
// fmt.Sprintf生成的字符串分配在堆上，对应数据地址可修改。


func Benchmark_Normal(b *testing.B) {
	b.N = N
	for i := 1; i < N; i++ {
		s := fmt.Sprintf("12345678901234567890123456789012345678901234567890")
		bb := []byte(s)
		bb[0] = 'x'
		s = string(bb)
		s = s
	}
}

// 只有动态生成的string，数据可以这样转换修改
func Benchmark_Direct(b *testing.B) {
	b.N = N
	for i := 1; i < N; i++ {
		s := fmt.Sprintf("12345678901234567890123456789012345678901234567890")
		bb := StringBytes(s)
		bb[0] = 'x'
		s = s
	}
}