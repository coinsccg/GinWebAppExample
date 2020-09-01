package main

import (
	"fmt"
	"unsafe"
)

// 从结构体字段排布看内存对齐

type T1 struct {
	A int8   // 1byte
	B string // 16byte
	C int8   // 1byte
}

type T2 struct {
	A int8
	C int8
	B string
}

func main() {
	v1 := T1{10, "你好", 20} // 32
	v2 := T2{10, 20, "你好"} // 24

	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(v2))
}
