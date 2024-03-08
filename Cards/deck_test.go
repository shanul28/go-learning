package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Not as Expected %v", len(cards))
	}
}

func TestSaveFile(t *testing.T) {
	filename := "_decktesting"
	err := os.Remove(filename)

	if err != nil {
		cards := newDeck()
		e := cards.saveToFile(filename)
		if e != nil {
			t.Errorf("Unable to save files %v", e)
		}
		cardsFF := newDeckFromFile(filename)
		if len(cardsFF) != len(cards) {
			t.Errorf("wrong no of deck found from file original: %v, fromFile: %v", len(cards), len(cardsFF))
		}
		os.Remove(filename)
	}

}
func TestDeal(t *testing.T) {
	cards := newDeck()
	handSize := 5
	handSizeCards, dealCards := deal(cards, handSize)

	if len(handSizeCards) != 5 && len(dealCards) != 9 {
		t.Errorf("Invalid Cards length%v%v", handSizeCards, dealCards)
	}
}
