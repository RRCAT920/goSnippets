package main

import (
	"errors"
	"fmt"
)

// 用于测试代码
// like
// import _ "package path"
func main() {
	foo()
}

func foo() {
	err := errors.New("jango")
	if err != nil {
		panic(err)
	}
	fmt.Println("jklasd")
}
