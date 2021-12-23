package bingoCard_test

import (
	"fmt"
	"testing"

	"github.com/maikeru/adventofcode2021/day_04_1/bingoCard"
)

func TestBingoCardWinsForARowOfDraws(t *testing.T) {
	cardData := []string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}
	card := bingoCard.New(cardData)
	if card.HasWon() {
		t.Errorf("card has won prematurely %v", card)
	}
	card.Draw(8)
	card.Draw(2)
	card.Draw(23)
	card.Draw(4)
	card.Draw(24)
	fmt.Println("card", card)
	if !card.HasWon() {
		t.Errorf("card should have won %v", card)
	}
}

func TestBingoCardWinsForAColumnOfDraws(t *testing.T) {
	cardData := []string{
		"22 13 17 11  0",
		" 8  2 23  4 24",
		"21  9 14 16  7",
		" 6 10  3 18  5",
		" 1 12 20 15 19",
	}
	card := bingoCard.New(cardData)
	if card.HasWon() {
		t.Errorf("card has won prematurely:\n%v", card)
	}
	card.Draw(17)
	card.Draw(23)
	card.Draw(14)
	card.Draw(3)
	card.Draw(20)
	fmt.Println("card", card)
	if !card.HasWon() {
		t.Errorf("card should have won:\n%v", card)
	}
}

func TestBingoCardScore(t *testing.T) {
	cardData := []string{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	}
	card := bingoCard.New(cardData)

	draws := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	lastDrawn := -1
	for i := range draws {
		draw := draws[i]
		card.Draw(draw)
		if card.HasWon() {
			lastDrawn = draw
			break
		}
	}

	expectedScore := 4512

	actualScore := card.Score(lastDrawn)
	if actualScore != expectedScore {
		t.Errorf("Incorrect score! Expected %d but got %d", expectedScore, actualScore)
	}
}
