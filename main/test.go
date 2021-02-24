package main

import (
	"fmt"
	_ "goSnippets/init1"
	_ "goSnippets/init2"
)

// 用于测试代码
func main() {
	fmt.Println(SafeFtoi(32.6))
}
