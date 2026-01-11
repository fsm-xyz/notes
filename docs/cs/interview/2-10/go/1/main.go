package main

import (
	"fmt"
	"sort"
	"strings"
)

type Segment struct {
	Position  int
	Intensity int
}

type IntensitySegments struct {
	segments []Segment
}

func NewIntensitySegments() *IntensitySegments {
	return &IntensitySegments{
		segments: make([]Segment, 0),
	}
}

func (is *IntensitySegments) Add(from, to, amount int) {
	if from >= to || amount == 0 {
		return
	}

	is.ensureSegmentAt(from)
	is.ensureSegmentAt(to)

	for i := range is.segments {
		pos := is.segments[i].Position
		if pos >= from && pos < to {
			is.segments[i].Intensity += amount
		}
	}

	is.cleanup()
}

func (is *IntensitySegments) Set(from, to, amount int) {
	if from >= to {
		return
	}

	intensityAtTo := is.getIntensityAt(to)

	newSegments := make([]Segment, 0)
	for _, seg := range is.segments {
		if seg.Position < from || seg.Position >= to {
			newSegments = append(newSegments, seg)
		}
	}
	is.segments = newSegments

	is.insertSegment(from, amount)

	if intensityAtTo != amount {
		is.insertSegment(to, intensityAtTo)
	}

	is.cleanup()
}

func (is *IntensitySegments) ToString() string {
	if len(is.segments) == 0 {
		return "[]"
	}

	result := make([]string, 0)

	for i, seg := range is.segments {
		if seg.Intensity != 0 {
			result = append(result, fmt.Sprintf("[%d,%d]", seg.Position, seg.Intensity))
		} else {
			hasNonZeroBefore := i > 0 && is.segments[i-1].Intensity != 0
			hasNonZeroAfter := false

			for j := i + 1; j < len(is.segments); j++ {
				if is.segments[j].Intensity != 0 {
					hasNonZeroAfter = true
					break
				}
			}

			if hasNonZeroBefore || hasNonZeroAfter {
				result = append(result, fmt.Sprintf("[%d,%d]", seg.Position, seg.Intensity))
			}
		}
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (is *IntensitySegments) ensureSegmentAt(pos int) {
	for _, seg := range is.segments {
		if seg.Position == pos {
			return
		}
	}

	intensity := is.getIntensityAt(pos)

	is.insertSegment(pos, intensity)
}

func (is *IntensitySegments) getIntensityAt(pos int) int {
	intensity := 0

	for _, seg := range is.segments {
		if seg.Position <= pos {
			intensity = seg.Intensity
		} else {
			break
		}
	}

	return intensity
}

func (is *IntensitySegments) insertSegment(pos, intensity int) {
	newSeg := Segment{Position: pos, Intensity: intensity}
	is.segments = append(is.segments, newSeg)

	sort.Slice(is.segments, func(i, j int) bool {
		return is.segments[i].Position < is.segments[j].Position
	})
}

func (is *IntensitySegments) cleanup() {
	if len(is.segments) <= 1 {
		return
	}

	cleaned := make([]Segment, 0, len(is.segments))
	cleaned = append(cleaned, is.segments[0])

	for i := 1; i < len(is.segments); i++ {
		current := is.segments[i]
		last := cleaned[len(cleaned)-1]

		if current.Intensity != last.Intensity {
			cleaned = append(cleaned, current)
		}
	}

	is.segments = cleaned
}

func main() {
	segments := NewIntensitySegments()

	fmt.Println("Initial:", segments.ToString()) // Should be "[]"

	segments.Add(10, 30, 1)
	fmt.Println("After add(10,30,1):", segments.ToString()) // Should be "[[10,1],[30,0]]"

	segments.Add(20, 40, 1)
	fmt.Println("After add(20,40,1):", segments.ToString()) // Should be "[[10,1],[20,2],[30,1],[40,0]]"

	segments.Add(10, 40, -2)
	fmt.Println("After add(10,40,-2):", segments.ToString()) // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"

	fmt.Println("----------------------------------------------")

	segments2 := NewIntensitySegments()
	segments2.Add(10, 30, 1)
	fmt.Println("After Add(10,30,1): set:", segments2.ToString()) // Should be "[[10,1],[30,0]]"

	segments2.Add(20, 40, 1)
	fmt.Println("After Add(20,40,1):", segments2.ToString()) // Should be "[[10,1],[20,2],[30,1],[40,0]]"
	segments2.Add(10, 40, -1)
	fmt.Println("After Add(10,40,-1):", segments2.ToString()) // Should be "[[20,1],[30,0]]"

	segments2.Add(10, 40, -1)
	fmt.Println("After Add(10,40,-1):", segments2.ToString()) // Should be "[[10,-1],[20,0],[30,-1],[40,0]]"
}
