package main

import (
	"fmt"
	"math"
)

// 安全的float64转成int32
func SafeFtoi(f float64) int32 {
	if math.MinInt32 <= f && f <= math.MaxInt32 {
		// why not int(f+0.5)?
		i, frac := math.Modf(f)
		if frac >= 0.5 {
			i++
		}
		return int32(i)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", f))
}
