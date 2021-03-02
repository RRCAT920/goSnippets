package main

import (
	"bufio"
	"fmt"
	"strings"
)

func ReaderDefaultBehavior() {
	str := strings.NewReader("bufio.Reader#ReadString will read the delim like \n")
	rd := bufio.NewReader(str)
	res, _ := rd.ReadString('\n')
	fmt.Printf("%q\n", res)
}

func init() {
	DefaultLogger.Log()
	ReaderDefaultBehavior()
}
