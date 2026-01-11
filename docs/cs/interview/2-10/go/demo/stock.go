package main

import "math"

func main() {
	arr := []float64{7, 1, 5, 3, 6, 4}
	trade(arr)
}

func trade(arr []float64) {
	profit := 0.0
	for i := 1; i < len(arr); i++ {
		tmp := arr[i] - arr[i-1]
		if tmp > 0 {
			profit += math.Max(0, tmp)
		}
	}
}
