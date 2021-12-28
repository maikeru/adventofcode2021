package diagram

import (
	"github.com/maikeru/adventofcode2021/day_05_1/lines"
)

type Diagram struct {
	lines []lines.Line
}

func New(lines []lines.Line) Diagram {
	return Diagram{lines: lines}
}

func (diagram *Diagram) GetOverlapCount() int {
	allPoints := make(map[string]int)
	for i := range diagram.lines {
		line := diagram.lines[i]
		linePoints := line.GetPoints()
		for j := range linePoints {
			linePoint := linePoints[j]
			value := allPoints[linePoint.String()]
			allPoints[linePoint.String()] = value + 1
		}
	}

	overlapCount := 0
	for _, value := range allPoints {
		if value > 1 {
			overlapCount++
		}
	}

	return overlapCount
}
