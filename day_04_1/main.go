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

func filter(cards *[]bingoCard.BingoCard, cardToFilter *bingoCard.BingoCard) []bingoCard.BingoCard {
	filteredCards := make([]bingoCard.BingoCard, 0)
	for i := range *cards {
		currentCard := &(*cards)[i]
		if currentCard != cardToFilter {
			filteredCards = append(filteredCards, *currentCard)
		}
	}
	return filteredCards
}

func Run() {
	// data := utils.LoadFile("./day_04_1/testData.txt")
	data := utils.LoadFile("./day_04_1/input.txt")

	bingoData := ParseBingoData(data)

	fmt.Printf("Bingo data!!! \n%v\n", bingoData)

	winningDraws := make([]int, 0)
	winningCards := make([]bingoCard.BingoCard, 0)

	remainingCards := bingoData.Cards

	for drawIndex := range bingoData.Draws {
		draw := bingoData.Draws[drawIndex]
		tempRemainingCards := make([]bingoCard.BingoCard, 0)
		fmt.Println("Drawn number:", draw)
		for cardIndex := range remainingCards {
			card := &remainingCards[cardIndex]
			card.Draw(draw)
			fmt.Printf("Card: \n%v\n", *card)
			if card.HasWon() {
				winningDraws = append(winningDraws, draw)
				winningCards = append(winningCards, *card)
			} else {
				// Keep checking non-winning cards in later draws
				tempRemainingCards = append(tempRemainingCards, *card)
			}
		}
		remainingCards = tempRemainingCards
	}

	if len(winningCards) == 0 {
		panic("Failed to find winning card")
	}

	firstWinDraw := winningDraws[0]
	firstWinCard := winningCards[0]

	score := firstWinCard.Score(firstWinDraw)

	fmt.Println("=== First winning card ===")
	fmt.Printf("Last drawn number: %d\n", firstWinDraw)
	fmt.Printf("Winning card: \n%v\n", firstWinCard)
	fmt.Printf("Score: %d\n", score)

	lastWinDraw := winningDraws[len(winningDraws)-1]
	lastWinCard := winningCards[len(winningCards)-1]
	lastWinScore := lastWinCard.Score(lastWinDraw)

	fmt.Println("=== Last winning card ===")
	fmt.Printf("Last drawn number: %d\n", lastWinDraw)
	fmt.Printf("Winning card: \n%v\n", lastWinCard)
	fmt.Printf("Score: %d\n", lastWinScore)
	return
}
