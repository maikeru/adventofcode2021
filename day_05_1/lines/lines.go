package lines

import (
	"github.com/maikeru/adventofcode2021/day_05_1/point"
	"github.com/maikeru/adventofcode2021/utils"
)

type Line struct {
	Start point.Point
	End   point.Point
}

func New(start point.Point, end point.Point) Line {
	return Line{Start: start, End: end}
}

func generateRange(start int, end int, size int) []int {
	generatedRange := make([]int, 0)
	if start < end {
		for i := start; i <= end; i++ {
			generatedRange = append(generatedRange, i)
		}
	} else if start > end {
		for i := start; i >= end; i-- {
			generatedRange = append(generatedRange, i)
		}
	} else {
		// If start and end are equal (i.e. 0, fill the range with start value up to size)
		for i := 0; i <= size; i++ {
			generatedRange = append(generatedRange, start)
		}
	}

	return generatedRange
}

func (line *Line) GetPoints() []point.Point {
	points := make([]point.Point, 0)
	sizeX := line.End.X - line.Start.X
	sizeY := line.End.Y - line.Start.Y
	var size int
	if sizeX == 0 {
		size = utils.Abs(sizeY)
	} else {
		size = utils.Abs(sizeX)
	}

	// It's diagonal!
	xRange := generateRange(line.Start.X, line.End.X, size)
	yRange := generateRange(line.Start.Y, line.End.Y, size)

	for i := range xRange {
		x := xRange[i]
		y := yRange[i]
		points = append(points, point.Point{X: x, Y: y})
	}

	return points
}

func (line *Line) IsDiagonal() bool {
	if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
		return false
	}
	return true
}
