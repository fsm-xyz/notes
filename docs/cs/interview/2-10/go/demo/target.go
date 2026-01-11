package main

import (
	"math"
	"slices"
)

func main() {

}

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。

func findClosest(arr []float64, target float64) float64 {
	slices.Sort(arr)
	res := arr[0] + arr[1] + arr[2]
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if math.Abs(target-sum) < math.Abs(target-res) {
				res = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return res
			}
		}
	}
	return res
}
