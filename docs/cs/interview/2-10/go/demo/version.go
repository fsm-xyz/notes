package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 给你两个版本号 `version1` 和 `version2` ，请你比较它们。
// 版本号由多个部分组成，每个部分由一个 '.' 连接。每个部分都由数字组成，可能包含前导零。
// 返回规则如下：
// 如果 version1 > version2 返回 1，
// 如果 version1 < version2 返回 -1，
// 除此之外返回 0

// 示例：
// 1.0 > 0.1
// 1.0 < 1.0.1 1.0 1.01
// 1.1 == 1.01
// 1.0.0 == 1
func main() {
	a := "1.0"
	b := "1.0.1"
	fmt.Println(compare(a, b))
}

func compare(a, b string) int {
	aa := strings.Split(a, ".")
	bb := strings.Split(b, ".")

	n := max(len(bb), len(aa))
	for i := 0; i < n; i++ {
		var aInt, bInt int
		if i < len(aa) {
			aInt, _ = strconv.Atoi(aa[i])
		}
		if i < len(bb) {
			bInt, _ = strconv.Atoi(bb[i])
		}
		if aInt > bInt {
			return 1
		} else if aInt < bInt {
			return -1
		}
	}
	return 0
}
