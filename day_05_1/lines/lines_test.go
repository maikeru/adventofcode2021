package lines_test

import (
	"reflect"
	"testing"

	"github.com/maikeru/adventofcode2021/day_05_1/lines"
)

func TestGeneratePointsOnLineWhenIncreasing(t *testing.T) {
	lineY := lines.New(lines.Point{0, 0}, lines.Point{0, 2})

	expectedPointsY := []lines.Point{
		{0, 0},
		{0, 1},
		{0, 2},
	}

	if !reflect.DeepEqual(lineY.GetPoints(), expectedPointsY) {
		t.Errorf("LineY drew incorrect points got %v but expected %v", lineY.GetPoints(), expectedPointsY)
	}

	lineX := lines.New(lines.Point{0, 0}, lines.Point{2, 0})

	expectedPointsX := []lines.Point{
		{0, 0},
		{1, 0},
		{2, 0},
	}

	if !reflect.DeepEqual(lineX.GetPoints(), expectedPointsX) {
		t.Errorf("LineX drew incorrect points got %v but expected %v", lineX.GetPoints(), expectedPointsX)
	}
}

func TestGeneratePointsOnLineWhenDecreasing(t *testing.T) {
	lineY := lines.New(lines.Point{0, 2}, lines.Point{0, 0})

	expectedPointsY := []lines.Point{
		{0, 2},
		{0, 1},
		{0, 0},
	}

	if !reflect.DeepEqual(lineY.GetPoints(), expectedPointsY) {
		t.Errorf("LineY drew incorrect points got %v but expected %v", lineY.GetPoints(), expectedPointsY)
	}

	lineX := lines.New(lines.Point{2, 0}, lines.Point{0, 0})

	expectedPointsX := []lines.Point{
		{2, 0},
		{1, 0},
		{0, 0},
	}

	if !reflect.DeepEqual(lineX.GetPoints(), expectedPointsX) {
		t.Errorf("LineX drew incorrect points got %v but expected %v", lineX.GetPoints(), expectedPointsX)
	}
}
