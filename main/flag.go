package main

// 表示资源的使用状态
const (
	Open = 1 << iota
	Close
	Pending
)
