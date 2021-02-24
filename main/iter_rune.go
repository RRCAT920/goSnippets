package main

import "fmt"

// 迭代Unicode字符串
func IterUnicodeStr(s string) {
	for _, char := range s {
		fmt.Printf("%c", char)
	}
}
