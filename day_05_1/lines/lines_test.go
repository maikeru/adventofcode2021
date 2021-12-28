package lines_test

import (
	"reflect"
	"testing"

	"github.com/maikeru/adventofcode2021/day_05_1/lines"
	"github.com/maikeru/adventofcode2021/day_05_1/point"
)

func TestGeneratePointsOnLineWhenIncreasing(t *testing.T) {
	lineY := lines.New(point.Point{X: 0, Y: 0}, point.Point{X: 0, Y: 2})

	expectedPointsY := []point.Point{
		{X: 0, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: 2},
	}

	if !reflect.DeepEqual(lineY.GetPoints(), expectedPointsY) {
		t.Errorf("LineY drew incorrect points got %v but expected %v", lineY.GetPoints(), expectedPointsY)
	}

	lineX := lines.New(point.Point{X: 0, Y: 0}, point.Point{X: 2, Y: 0})

	expectedPointsX := []point.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
	}

	if !reflect.DeepEqual(lineX.GetPoints(), expectedPointsX) {
		t.Errorf("LineX drew incorrect points got %v but expected %v", lineX.GetPoints(), expectedPointsX)
	}
}

func TestGeneratePointsOnLineWhenDecreasing(t *testing.T) {
	lineY := lines.New(point.Point{X: 0, Y: 2}, point.Point{X: 0, Y: 0})

	expectedPointsY := []point.Point{
		{X: 0, Y: 2},
		{X: 0, Y: 1},
		{X: 0, Y: 0},
	}

	if !reflect.DeepEqual(lineY.GetPoints(), expectedPointsY) {
		t.Errorf("LineY drew incorrect points got %v but expected %v", lineY.GetPoints(), expectedPointsY)
	}

	lineX := lines.New(point.Point{X: 2, Y: 0}, point.Point{X: 0, Y: 0})

	expectedPointsX := []point.Point{
		{X: 2, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 0},
	}

	if !reflect.DeepEqual(lineX.GetPoints(), expectedPointsX) {
		t.Errorf("LineX drew incorrect points got %v but expected %v", lineX.GetPoints(), expectedPointsX)
	}
}
