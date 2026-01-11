package main

import "fmt"

// [
//   ("Hello", 0, 4),
//   ("World", 5, 10),
//   ("This", 13, 18),
//   ("is", 18, 20),
//   ("a", 20, 21),
//   ("test", 21, 23)
// ]
// 字幕长度：12

// 输出：
// [
//
//	("Hello World", 0, 10),
//	("This is a", 13, 21),
//	("test", 21, 23)
//
// ]

// Subtitle 代表一个字幕片段
type Subtitle struct {
	Text  string
	Start int
	End   int
}

// mergeSubtitles 将字幕片段按照最大长度合并
func mergeSubtitles(subtitles []Subtitle, maxLength int) []Subtitle {
	if len(subtitles) == 0 {
		return []Subtitle{}
	}

	var result []Subtitle
	currentText := ""
	currentStart := 0
	currentEnd := 0

	for _, subtitle := range subtitles {
		var candidateText string

		if currentText == "" {
			// 第一个字幕片段
			candidateText = subtitle.Text
			currentStart = subtitle.Start
		} else {
			// 尝试添加新的字幕片段
			candidateText = currentText + " " + subtitle.Text
		}

		if len(candidateText) <= maxLength {
			// 可以合并
			currentText = candidateText
			currentEnd = subtitle.End
			continue
		}

		result = append(result, Subtitle{
			Text:  currentText,
			Start: currentStart,
			End:   currentEnd,
		})

		// 开始新的合并
		currentText = subtitle.Text
		currentStart = subtitle.Start
		currentEnd = subtitle.End
	}

	// 添加最后一个合并的结果
	if currentText != "" {
		result = append(result, Subtitle{
			Text:  currentText,
			Start: currentStart,
			End:   currentEnd,
		})
	}
	return result
}

func main() {
	subtitles := []Subtitle{
		{"Hello", 0, 4},
		{"World", 5, 10},
		{"This", 13, 18},
		{"is", 18, 20},
		{"a", 20, 21},
		{"test", 21, 23},
	}

	maxLength := 12
	merged := mergeSubtitles(subtitles, maxLength)

	fmt.Println("原始字幕:")
	for _, sub := range subtitles {
		fmt.Printf("(\"%s\", %d, %d)\n", sub.Text, sub.Start, sub.End)
	}

	fmt.Printf("\n最大长度: %d\n\n", maxLength)

	fmt.Println("合并后的字幕:")
	for _, sub := range merged {
		fmt.Printf("(\"%s\", %d, %d)\n", sub.Text, sub.Start, sub.End)
	}
}
