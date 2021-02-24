package main

import "math/big"

// 阶乘算法
func fac(n int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		res.Mul(res, big.NewInt(i))
	}
	return res
}
