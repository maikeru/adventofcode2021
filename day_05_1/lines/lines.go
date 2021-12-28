package lines

import (
	"fmt"

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
	sizeY := line.End.Y - line.Start.Y

	if sizeX == 0 {
		yRange := generateRange(line.Start.Y, line.End.Y)
		for i := range yRange {
			y := yRange[i]
			points = append(points, point.Point{line.Start.X, y})
		}
	} else if sizeY == 0 {
		xRange := generateRange(line.Start.X, line.End.X)
		for i := range xRange {
			x := xRange[i]
			points = append(points, point.Point{x, line.Start.Y})
		}
	} else {
		// It's diagonal!
		xRange := generateRange(line.Start.X, line.End.X)
		yRange := generateRange(line.Start.Y, line.End.Y)

		if utils.Abs(sizeX) != utils.Abs(sizeY) {
			panic(fmt.Sprintf("Not 45 deg line! sizeX: %d sizeY: %d", sizeX, sizeY))
		}

		for i := range xRange {
			x := xRange[i]
			y := yRange[i]
			points = append(points, point.Point{x, y})
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
