package day_05_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/day_05_1/diagram"
	"github.com/maikeru/adventofcode2021/day_05_1/lines"
	"github.com/maikeru/adventofcode2021/day_05_1/point"
	"github.com/maikeru/adventofcode2021/utils"
)

func endToPoint(end string) point.Point {
	coords := strings.Split(end, ",")
	x := utils.StringToInt(coords[0])
	y := utils.StringToInt(coords[1])
	return point.Point{X: x, Y: y}
}

func StringToLine(entry string) lines.Line {
	ends := strings.Split(entry, " -> ")
	start := endToPoint(ends[0])
	end := endToPoint(ends[1])
	return lines.New(start, end)
}

func StringsToLines(entries []string) []lines.Line {
	lines := make([]lines.Line, 0)
	for i := range entries {
		entry := entries[i]
		line := StringToLine(entry)
		if !line.IsDiagonal() {
			lines = append(lines, line)
		}
	}
	return lines
}

func GetOverlapCount(lines []lines.Line) int {
	// Create diagram with lines on it
	diagram := diagram.New(lines)

	// Return count of overlaps
	return diagram.GetOverlapCount()
}

func Run() {
	data := utils.LoadFile("./day_05_1/testInput.txt")
	// data := utils.LoadFile("./day_05_1/input.txt")
	lines := StringsToLines(data)
	overlapCount := GetOverlapCount(lines)
	fmt.Println("Overlap count:", overlapCount)
}
