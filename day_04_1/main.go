package day_04_1

import (
	"fmt"
	"strings"

	"github.com/maikeru/adventofcode2021/day_04_1/bingoCard"
	"github.com/maikeru/adventofcode2021/utils"
)

type bingoData struct {
	Draws []int
	Cards []bingoCard.BingoCard
}

func ParseBingoData(data []string) bingoData {
	draws := utils.MapStringsToInts(strings.Split(data[0], ","))
	remainingData := data[2:]

	currentCardData := make([]string, 0)
	cards := make([]bingoCard.BingoCard, 0)
	for lineIndex := range remainingData {
		line := remainingData[lineIndex]
		fmt.Println("line:", line)
		if line == "" {
			// Blank line: we should have enough data to create a card
			card := bingoCard.New(currentCardData)
			cards = append(cards, card)

			// Clear card data ready for next card
			currentCardData = make([]string, 0)

			// Skip the blank line
			continue
		}

		currentCardData = append(currentCardData, line)
	}

	result := bingoData{Draws: draws, Cards: cards}

	return result
}

func Run() {
	// data := utils.LoadFile("./day_04_1/testData.txt")
	data := utils.LoadFile("./day_04_1/input.txt")

	bingoData := ParseBingoData(data)

	fmt.Printf("Bingo data!!! \n%v\n", bingoData)

	lastDrawn := -1
	var winningCard bingoCard.BingoCard
	for drawIndex := range bingoData.Draws {
		draw := bingoData.Draws[drawIndex]
		for cardIndex := range bingoData.Cards {
			card := &bingoData.Cards[cardIndex]
			fmt.Println("Drawn number:", draw)
			card.Draw(draw)
			fmt.Printf("Card: \n%v\n", *card)
			if card.HasWon() {
				lastDrawn = draw
				winningCard = *card
				score := winningCard.Score(lastDrawn)

				fmt.Printf("Last drawn number: %d\n", lastDrawn)
				fmt.Printf("Winning card: \n%v\n", winningCard)
				fmt.Printf("Score: %d\n", score)
				return
			}
		}
	}

	panic("Failed to find winning card")
}
