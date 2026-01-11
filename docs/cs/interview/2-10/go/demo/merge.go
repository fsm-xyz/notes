package main

func  main() {

}

func merge(a []int, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	i, j := 0, 0
	res := make([]int, 0, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	if i < len(a) {
		res = append(res, a[i:]...)
		i++
	}

	if j < len(b) {
		res = append(res, b[j:]...)
		j++
	}
	return res
}
