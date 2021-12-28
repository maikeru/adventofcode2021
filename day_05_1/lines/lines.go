package lines

import (
	"github.com/maikeru/adventofcode2021/day_05_1/point"
)

type Line struct {
	Start point.Point
	End   point.Point
}

func New(start point.Point, end point.Point) Line {
	return Line{Start: start, End: end}
}

func generateRange(start int, end int) []int {
	generatedRange := make([]int, 0)
	if start < end {
		for i := start; i <= end; i++ {
			generatedRange = append(generatedRange, i)
		}
	} else {
		for i := start; i >= end; i-- {
			generatedRange = append(generatedRange, i)
		}
	}

	return generatedRange
}

func (line *Line) GetPoints() []point.Point {
	points := make([]point.Point, 0)
	sizeX := line.End.X - line.Start.X

	if sizeX == 0 {
		yRange := generateRange(line.Start.Y, line.End.Y)
		for i := range yRange {
			y := yRange[i]
			points = append(points, point.Point{line.Start.X, y})
		}
	} else {
		xRange := generateRange(line.Start.X, line.End.X)
		for i := range xRange {
			x := xRange[i]
			points = append(points, point.Point{x, line.Start.Y})
		}
	}

	return points
}

func (line *Line) IsDiagonal() bool {
	if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
		return false
	}
	return true
}
