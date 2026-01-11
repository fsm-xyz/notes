package main

func main() {

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func find(arr []int) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	}
	n := len(arr)
	if arr[0] >= 0 {
		return arr[0], true
	}
	if arr[n-1] <= 0 {
		return arr[n-1], true
	}

	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if arr[mid] < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 {
		return arr[0], true
	}

	if abs(arr[left-1]) <= abs(arr[left]) {
		return arr[left-1], true
	} else {
		return arr[left], true
	}
}
