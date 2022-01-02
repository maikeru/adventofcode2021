package day_08_1_test

import (
	"reflect"
	"testing"

	"github.com/maikeru/adventofcode2021/day_08_1"
	"github.com/maikeru/adventofcode2021/utils"
)

func TestCountEasyNumbersInOutput(t *testing.T) {
	data := utils.LoadFile("./testInput.txt")
	outputs := day_08_1.GetOutputs(data)
	count := day_08_1.CountEasyNumbersInOutput(outputs)
	expectedCount := 26

	if count != expectedCount {
		t.Errorf("Incorrect count! Expected %d but got %d", expectedCount, count)
	}
}

func TestFindPatternsWithLength(t *testing.T) {
	patterns := []string{"abc", "bc", "defg", "b"}
	expectedPatterns := []string{"bc"}
	actualPatterns := day_08_1.FindPatternsWithLength(patterns, 2)
	if !reflect.DeepEqual(actualPatterns, expectedPatterns) {
		t.Errorf("Did not find patterns with length 2! Expected %v but got %v", expectedPatterns, actualPatterns)
	}
}

func TestFindPatternsContaining(t *testing.T) {
	patterns := []string{"abc", "bc", "defg", "b"}
	search := "bc"
	expectedPatterns := []string{"abc", "bc"}
	actualPatterns := day_08_1.FindPatternsContaining(patterns, search)
	if !reflect.DeepEqual(expectedPatterns, actualPatterns) {
		t.Errorf("Incorrect patterns! Expected %v but got %v", expectedPatterns, actualPatterns)
	}
}

func TestFindPatternsNotContaining(t *testing.T) {
	patterns := []string{"abc", "bc", "defg", "b"}
	search := "bc"
	expectedPatterns := []string{"defg", "b"}
	actualPatterns := day_08_1.FindPatternsNotContaining(patterns, search)
	if !reflect.DeepEqual(expectedPatterns, actualPatterns) {
		t.Errorf("Incorrect patterns! Expected %v but got %v", expectedPatterns, actualPatterns)
	}
}

func TestFilterWholePatterns(t *testing.T) {
	patterns := []string{"abc", "bc", "defg", "b"}
	search := []string{"bc", "defg"}
	expectedPatterns := []string{"abc", "b"}
	actualPatterns := day_08_1.FilterWholePatterns(patterns, search)
	if !reflect.DeepEqual(expectedPatterns, actualPatterns) {
		t.Errorf("Incorrect patterns! Expected %v but got %v", expectedPatterns, actualPatterns)
	}
}

func TestDoesPatternMatchWhenShouldMatch(t *testing.T) {
	pattern := "cba"
	toMatch := "abc"
	if !day_08_1.DoesPatternMatch(pattern, toMatch) {
		t.Errorf("Pattern should match but didn't! pattern: %s toMatch: %s", pattern, toMatch)
	}
}

func TestDoesPatternMatchWhenInputsAreDifferentLengths(t *testing.T) {
	pattern := "cb"
	toMatch := "abc"
	if day_08_1.DoesPatternMatch(pattern, toMatch) {
		t.Errorf("Pattern should not match but did! pattern: %s toMatch: %s", pattern, toMatch)
	}
}

func TestDoesPatternMatchWhenInputsAreSameLengths(t *testing.T) {
	pattern := "cbe"
	toMatch := "abc"
	if day_08_1.DoesPatternMatch(pattern, toMatch) {
		t.Errorf("Pattern should not match but did! pattern: %s toMatch: %s", pattern, toMatch)
	}
}

func TestDecodeLine(t *testing.T) {
	line := day_08_1.Entry{
		UniquePatterns: []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		Outputs:        []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
	}

	expectedOutput := 5353
	output := day_08_1.DecodeLine(line)

	if output != expectedOutput {
		t.Errorf("Incorrect output! Expected [%d] but got [%d]", expectedOutput, output)
	}

}
