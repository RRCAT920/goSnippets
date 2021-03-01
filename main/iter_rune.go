package main

import "fmt"

// 迭代Unicode字符串
func IterUnicodeStr(s string) {
	for _, char := range s {
		fmt.Printf("%c", char)
	}
	fmt.Println()
}

func init() {
	DefaultLogger.Log()
	IterUnicodeStr("你好中国")
}
