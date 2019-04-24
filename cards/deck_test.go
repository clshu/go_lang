package main

import (
	"os"
	"testing"
)

const CardNumber = 16
const FirstElement = "Ace of Spades"
const LastElement = "Four of Clubs"
const CardFile = "go_cards"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != CardNumber {
		t.Errorf("Expected length of %d but get %d", CardNumber, len(d))
	}

	if d[0] != FirstElement {
		t.Errorf("Expected %s but get %s", FirstElement, d[0])
	}

	if d[len(d)-1] != LastElement {
		t.Errorf("Expected %s but get %s", LastElement, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove(CardFile)

	deck := newDeck()
	deck.saveToFile(CardFile)

	loadedDeck := newDeckFromFile(CardFile)
	if len(deck) != len(loadedDeck) {
		t.Errorf("Expected deck size %d but get %d", len(deck), len(loadedDeck))
	}

	os.Remove(CardFile)
}
