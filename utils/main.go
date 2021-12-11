package utils

import "strconv"

func MapStringsToInts(stringArray []string) []int {
	numArray := make([]int, len(stringArray))
	for i := range stringArray {
		num, err := strconv.Atoi(stringArray[i])
		if err != nil {
			panic(err)
		}
		numArray[i] = num
	}
	return numArray
}
