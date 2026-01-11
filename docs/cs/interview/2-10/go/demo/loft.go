package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(loft(40), time.Since(t))
	t1 := time.Now()
	fmt.Println(loft1(80), time.Since(t1))
	t2 := time.Now()
	fmt.Println(loft2(80, map[int]int{}), time.Since(t2))
	t3 := time.Now()
	fmt.Println(loft3(80), time.Since(t3))
}

// 简单递归
func loft(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	return loft(n-1) + loft(n-3)
}

// 递归转迭代 用map存储进行优化
var gm = map[int]int{}

func loft1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if _, ok := gm[n]; ok {
		return gm[n]
	}
	gm[n] = loft1(n-1) + loft1(n-3)
	return gm[n]
}

// 局部map
func loft2(n int, m map[int]int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = loft2(n-1, m) + loft2(n-3, m)
	return m[n]
}

// 迭代
func loft3(n int) int {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
	}
	if k, ok := m[n]; ok {
		return k
	}
	for i := 4; i <= n; i++ {
		m[i] = m[i-1] + m[i-3]
	}
	return m[n]
}
