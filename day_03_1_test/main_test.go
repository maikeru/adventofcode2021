package day_03_1_test

import (
	"testing"

	"github.com/maikeru/adventofcode2021/day_03_1"
)

func TestGetGammaAndEpsilon(t *testing.T) {
	gamma, epsilon := day_03_1.GetGammaAndEpsilon([]int{1, 1, 1, 1, 1})
	gammaWant := 31
	if gamma != gammaWant {
		t.Errorf("gamma should be %d but was %d", gammaWant, gamma)
	}

	epsilonWant := 0
	if epsilon != epsilonWant {
		t.Errorf("epsilon should be %d but was %d", epsilonWant, epsilon)
	}
}

func TestGetGammaAndEpsilonWithBigNumber(t *testing.T) {
	gamma, epsilon := day_03_1.GetGammaAndEpsilon([]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1})
	gammaWant := 31
	if gamma != gammaWant {
		t.Errorf("gamma should be %d but was %d", gammaWant, gamma)
	}

	epsilonWant := 4064
	if epsilon != epsilonWant {
		t.Errorf("epsilon should be %d but was %d", epsilonWant, epsilon)
	}
}
