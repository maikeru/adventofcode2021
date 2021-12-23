package bingoCard

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/utils"
)

type cardEntry struct {
	value   int
	isDrawn bool
}

type BingoCard struct {
	data [][]cardEntry
}

func MapStringsToEntries(stringArray []string) []cardEntry {
	entryArray := make([]cardEntry, len(stringArray))
	for i := range stringArray {
		entry := NewEntry(stringArray[i])
		entryArray[i] = entry
	}
	return entryArray
}

func NewEntry(value string) cardEntry {
	return cardEntry{utils.StringToInt(value), false}
}

func New(cardData []string) BingoCard {
	finalData := make([][]cardEntry, 0)
	for i := 0; i < len(cardData); i++ {
		line := cardData[i]
		entries := strings.Fields(line)
		finalData = append(finalData, MapStringsToEntries(entries))
	}
	return BingoCard{finalData}
}

func allEntriesDrawn(rowOrColumn *[]cardEntry) bool {
	fmt.Println("rowOrColumn", *rowOrColumn)
	for i := range *rowOrColumn {
		entry := (*rowOrColumn)[i]
		fmt.Println("entry", entry)
		if !entry.isDrawn {
			return false
		}
	}
	return true
}

func (card *BingoCard) HasWon() bool {
	for i := range card.data {
		row := &card.data[i]
		if allEntriesDrawn(row) {
			return true
		}
	}

	for i := range card.data[0] {
		column := make([]cardEntry, len(card.data))
		for j := range card.data {
			column[j] = card.data[j][i]
		}
		if allEntriesDrawn(&column) {
			return true
		}
	}
	return false
}

func (card *BingoCard) Draw(num int) {
	for i := range card.data {
		row := &card.data[i]
		for j := range *row {
			entry := &(*row)[j]
			if entry.value == num {
				entry.isDrawn = true
			}
		}
	}
	return
}

func (card *BingoCard) getUndrawn() []int {
	undrawn := make([]int, 0)
	for i := range card.data {
		row := card.data[i]
		for j := range row {
			if !row[j].isDrawn {
				undrawn = append(undrawn, row[j].value)
			}
		}
	}
	return undrawn
}

func (card *BingoCard) Score(lastDrawn int) int {
	sumOfUndrawn := utils.SumAll(card.getUndrawn())
	return sumOfUndrawn * lastDrawn
}

var colorReset = "\033[0m"

var colorRed = "\033[31m"

func (card BingoCard) String() string {
	output := ""
	for i := 0; i < len(card.data); i++ {
		row := card.data[i]
		rowOutput := ""
		for j := 0; j < len(row); j++ {
			entry := row[j]
			if entry.isDrawn {
				rowOutput = fmt.Sprintf("%s %s%2d%s", rowOutput, colorRed, entry.value, colorReset)
			} else {
				rowOutput = fmt.Sprintf("%s %2d", rowOutput, entry.value)
			}
		}
		output = fmt.Sprintf("%s %s\n", output, rowOutput)
	}
	return output
}
