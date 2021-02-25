package main

import (
	"fmt"
	"os"
	"runtime"
)

// 变量经典案例
func variable() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

func init() {
	DefaultLogger.Log()
	variable()
}
