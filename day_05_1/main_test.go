package day_05_1_test

import (
	"testing"

	day051 "github.com/maikeru/adventofcode2021/day_05_1"
	"github.com/maikeru/adventofcode2021/day_05_1/lines"
	"github.com/maikeru/adventofcode2021/day_05_1/point"
	"github.com/maikeru/adventofcode2021/utils"
)

func TestGetOverlappingLineCountWithoutDiagonals(t *testing.T) {
	data := utils.LoadFile("./testInput.txt")
	lines := day051.StringsToLines(data, true)
	count := day051.GetOverlapCount(lines)
	expectedCount := 5
	if count != expectedCount {
		t.Errorf("Incorrect count! Got %d but expected %d", count, expectedCount)
	}
}

func TestGetOverlappingLineCountWithDiagonals(t *testing.T) {
	data := utils.LoadFile("./testInput.txt")
	lines := day051.StringsToLines(data, false)
	count := day051.GetOverlapCount(lines)
	expectedCount := 12
	if count != expectedCount {
		t.Errorf("Incorrect count! Got %d but expected %d", count, expectedCount)
	}
}

func TestStringToLine(t *testing.T) {
	line := day051.StringToLine("0,9 -> 5,9")

	expectedLine := lines.New(point.Point{X: 0, Y: 9}, point.Point{X: 5, Y: 9})

	if line.Start != expectedLine.Start || line.End != expectedLine.End {
		t.Errorf("Incorrect line! Got %v but expected %v", line, expectedLine)
	}
}
