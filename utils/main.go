package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func StringToInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return num
}

func MapStringsToInts(stringArray []string) []int {
	numArray := make([]int, len(stringArray))
	for i := range stringArray {
		num := StringToInt(stringArray[i])
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
