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
		// fmt.Println("line start:", line.Start, "end:", line.End, "linePoints: ", linePoints)
		for j := range linePoints {
			linePoint := linePoints[j]
			value := allPoints[linePoint.String()]
			// fmt.Println("linePoint", linePoint, "value:", value+1)
			allPoints[linePoint.String()] = value + 1
		}
	}

	overlapCount := 0
	for _, value := range allPoints {
		if value > 1 {
			overlapCount++
		}
	}

	// debugAllPoints, err := json.Marshal(allPoints)
	// if err != nil {
	// 	panic(err)
	// }
	// println("allPoints", string(debugAllPoints))
	// println("length allPoints", len(allPoints))
	return overlapCount
}
