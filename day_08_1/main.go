package day_08_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

type Entry struct {
	UniquePatterns []string
	Outputs        []string
}

func isEasyNumber(entry string) bool {
	size := len(entry)
	switch size {
	case 2, 4, 3, 7: // 1, 4, 7, 8
		return true
	}
	return false
}

func CountEasyNumbersInOutput(outputs [][]string) int {
	total := 0
	for i := range outputs {
		outputLine := outputs[i]
		for j := range outputLine {
			output := outputLine[j]
			if isEasyNumber(output) {
				total++
			}
		}
	}
	return total
}

func GetOutputs(data []string) [][]string {
	outputs := make([][]string, 0)
	for i := range data {
		line := data[i]
		inputsAndOutput := strings.Split(line, " | ")
		output := strings.Fields(inputsAndOutput[1])
		outputs = append(outputs, output)
	}
	return outputs
}

type filterFunction func(item string) bool

func filterStrings(input []string, filterFunc filterFunction) []string {
	filtered := make([]string, 0)
	for i := range input {
		item := input[i]
		if filterFunc(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func FindPatternsWithLength(patterns []string, length int) []string {
	return filterStrings(patterns, func(item string) bool { return len(item) == length })
}

func findPatternDifference(p1 string, p2 string) []string {
	longestPattern := p1
	shortestPattern := p2
	if len(p2) > len(p1) {
		longestPattern = p2
		shortestPattern = p1
	}

	difference := make([]string, 0)
	for i := range longestPattern {
		if !strings.Contains(shortestPattern, string(longestPattern[i])) {
			difference = append(difference, string(longestPattern[i]))
		}
	}
	return difference
}

/*
Filters out patterns that don't contain everything in the search pattern (no partial matches)
*/
func FindPatternsContaining(patterns []string, search string) []string {
	return filterStrings(patterns, func(pattern string) bool {
		searchChars := strings.Split(search, "")
		for j := range searchChars {
			searchChar := searchChars[j]
			if !strings.Contains(pattern, searchChar) {
				return false
			}
		}

		return true
	})
}

/*
 Filters out patterns that contain everything in the search pattern (no partial matches)
*/
func FindPatternsNotContaining(patterns []string, search string) []string {
	return filterStrings(patterns, func(pattern string) bool {
		searchChars := strings.Split(search, "")
		// want to only match when the whole pattern matches
		matchCount := 0
		for j := range searchChars {
			searchChar := searchChars[j]
			if strings.Contains(pattern, searchChar) {
				matchCount++
			}
		}
		return matchCount < len(searchChars)
	})
}

/*
 Filter out patterns in the filter array from the input array. Exact matches only
*/
func FilterWholePatterns(input []string, filter []string) []string {
	return filterStrings(input, func(pattern string) bool {
		matches := false
		for i := range filter {
			filterPattern := filter[i]
			if pattern == filterPattern {
				matches = true
			}
		}
		return !matches
	})
}

func generateDecodeMap(uniquePatterns []string) map[string]string {
	decodeMap := make(map[string]string)

	sevenPattern := FindPatternsWithLength(uniquePatterns, 3)[0]
	decodeMap[sevenPattern] = "7"
	onePattern := FindPatternsWithLength(uniquePatterns, 2)[0]
	decodeMap[onePattern] = "1"
	fourPattern := FindPatternsWithLength(uniquePatterns, 4)[0]
	decodeMap[fourPattern] = "4"
	eightPattern := FindPatternsWithLength(uniquePatterns, 7)[0]
	decodeMap[eightPattern] = "8"

	// All the patterns with five segments
	twoThreeFivePatterns := FindPatternsWithLength(uniquePatterns, 5)

	// Only three contains both of the segments that exist in one
	threePattern := FindPatternsContaining(twoThreeFivePatterns, onePattern)[0]
	decodeMap[threePattern] = "3"

	// Subtracting the pattern for four from eight gives us a "bottom-left" pattern
	// That pattern exists in the pattern for two but not for five so should let us
	// distinguish the remaining numbers
	bottomLeftPattern := strings.Join(findPatternDifference(eightPattern, fourPattern), "")
	twoPattern := FindPatternsContaining(twoThreeFivePatterns, bottomLeftPattern)[0]
	decodeMap[twoPattern] = "2"
	fivePattern := FilterWholePatterns(twoThreeFivePatterns, []string{twoPattern, threePattern})[0]
	decodeMap[fivePattern] = "5"

	// All the patterns with six segments
	zeroSixNinePatterns := FindPatternsWithLength(uniquePatterns, 6)
	// Only six does not contain the segments that exist in one out of zero, six and nine
	sixPattern := FindPatternsNotContaining(zeroSixNinePatterns, onePattern)[0]
	decodeMap[sixPattern] = "6"
	// Nine contains three
	ninePattern := FindPatternsContaining(zeroSixNinePatterns, threePattern)[0]
	decodeMap[ninePattern] = "9"
	zeroPattern := FilterWholePatterns(zeroSixNinePatterns, []string{sixPattern, ninePattern})[0]
	decodeMap[zeroPattern] = "0"
	return decodeMap
}

/*
 Check if the two strings have the same characters but not necessarily in the same order
*/
func DoesPatternMatch(pattern string, toMatch string) bool {
	if len(pattern) != len(toMatch) {
		return false
	}

	for i := range toMatch {
		characterToMatch := string(toMatch[i])
		if !strings.Contains(pattern, characterToMatch) {
			return false
		}
	}
	return true
}

func decodeItem(decodeMap map[string]string, item string) string {
	for pattern, value := range decodeMap {
		if DoesPatternMatch(pattern, item) {
			return value
		}
	}
	panic(fmt.Sprintf("Unable to decode item %s", item))
}

func DecodeLine(entry Entry) int {
	decodeMap := generateDecodeMap(entry.UniquePatterns)
	decodedOutputs := ""
	for i := range entry.Outputs {
		output := entry.Outputs[i]
		decodedOutputs += decodeItem(decodeMap, output)
	}
	return utils.StringToInt(decodedOutputs)
}

func toEntry(line string) Entry {
	patternsAndOutputs := strings.Split(line, " | ")
	uniquePatterns := strings.Split(patternsAndOutputs[0], " ")
	outputs := strings.Split(patternsAndOutputs[1], " ")
	return Entry{UniquePatterns: uniquePatterns, Outputs: outputs}
}

func DecodeAndSumLines(data []string) int {
	total := 0
	for i := range data {
		line := data[i]
		entry := toEntry(line)
		value := DecodeLine(entry)
		total += value
	}
	return total
}

func Run() {
	// data := utils.LoadFile("./day_08_1/testInput.txt")
	data := utils.LoadFile("./day_08_1/input.txt")
	outputs := GetOutputs(data)
	count := CountEasyNumbersInOutput(outputs)
	fmt.Println("PART 1: count of easy numbers", count)
	total := DecodeAndSumLines(data)
	fmt.Println("PART 2: total of all outputs", total)
}
