package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

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

func LoadFile(filename string) []string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes[:])
	return strings.Split(fileString, "\n")
}
