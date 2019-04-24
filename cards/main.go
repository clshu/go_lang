package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()

	cards.print()
	fmt.Println("=============")

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	fmt.Println("=============")

	cards.saveToFile("card_deck")

	newCards := newDeckFromFile("card_deck")

	newCards.print()
	fmt.Println("=============")
	newCards.shuffle()
	newCards.print()
	fmt.Println("=============")
	newCards.shuffle()
	newCards.print()
	fmt.Println("=============")

	os.Remove("card_deck")
	even_or_odd()
}
