package day_03_2_test

import (
	"testing"

	"github.com/maikeru/adventofcode2021/day_03_2"
)

var reportTest = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	rating := day_03_2.GetOxygenGeneratorRating(reportTest)
	ratingWant := 23
	if rating != ratingWant {
		t.Errorf("rating should be %d but was %d", ratingWant, rating)
	}
}

func TestGetCO2ScrubberRating(t *testing.T) {
	rating := day_03_2.GetCO2ScrubberRating(reportTest)
	ratingWant := 10
	if rating != ratingWant {
		t.Errorf("rating should be %d but was %d", ratingWant, rating)
	}
}
