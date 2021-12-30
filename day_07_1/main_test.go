package day_07_1_test

import (
	"testing"

	"github.com/maikeru/adventofcode2021/day_07_1"
)

func TestGetFuelUsedForOptimumPositionPart1(t *testing.T) {
	startPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expectedFuel := 37
	actualFuel := day_07_1.GetFuelUsedForOptimumPositionPart1(startPositions)

	if expectedFuel != actualFuel {
		t.Errorf("Incorrect fuel used! Expected %d but got %d", expectedFuel, actualFuel)
	}
}
