package main

import (
	"fmt"
	"unsafe"
)

func NilSlice()  {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)
	s4 := *new([]int)

	var a1 = *(*[1]int)(unsafe.Pointer(&s1))// nil切片
	var a2 = *(*[1]int)(unsafe.Pointer(&s2))// 空切片
	var a3 = *(*[1]int)(unsafe.Pointer(&s3))
	var a4 = *(*[1]int)(unsafe.Pointer(&s4))
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}
