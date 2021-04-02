package main

import (
	"fmt"
	"unsafe"
)

const (
	q = "abc"
	w = len(q)
	e = unsafe.Sizeof(q)
)

func main() {
	//常量
	const LENTH int = 10
	const WIDTH int = 5

	var area int
	const a, b, c = 1, false, "str"
	area = LENTH * WIDTH
	fmt.Print("面积为：", area)

	println()
	println(a, b, c)
	println(q, w, e)
}
