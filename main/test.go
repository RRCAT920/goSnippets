package main

import "fmt"

type T func(int2 int)

// 用于测试代码
func main() {
	var a interface{}
	fmt.Println(a)
	var f T
	fmt.Println(f)
	IterUnicodeStr("Chinese: 日本語")
	print()
}
