package day_06_1_test

import (
	"reflect"
	"testing"

	"github.com/maikeru/adventofcode2021/day_06_1"
)

func TestGetFishCount(t *testing.T) {
	initialFish := []int{3, 4, 3, 1, 2}
	days := 18
	expectedFishCount := 26
	finalFishCount := day_06_1.GetFishCount(initialFish, days)

	if finalFishCount != expectedFishCount {
		t.Errorf("Incorrect fish count after %d days: got %d but expected %d", days, finalFishCount, expectedFishCount)
	}
}

func TestSimulateDayWithoutBirths(t *testing.T) {
	initialFish := []int{3, 4, 3, 1, 2}
	expectedFish := []int{2, 3, 2, 0, 1}
	actualFish := day_06_1.SimulateDay(&initialFish)

	if !reflect.DeepEqual(expectedFish, *actualFish) {
		t.Errorf("Incorrect fish! Expected %v but was %v", expectedFish, actualFish)
	}
}

func TestSimulateDayWithBirths(t *testing.T) {
	initialFish := []int{2, 3, 2, 0, 1}
	expectedFish := []int{1, 2, 1, 6, 0, 8}
	actualFish := day_06_1.SimulateDay(&initialFish)

	if !reflect.DeepEqual(expectedFish, *actualFish) {
		t.Errorf("Incorrect fish! Expected %v but was %v", expectedFish, actualFish)
	}
}
