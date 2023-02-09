package main

import (
	"fmt"
	"os"
)

func main() {
	deck := NewPlayingCardDeck()

	card := deck.RandomCard()
	fmt.Printf("card: %s\n", card)

	playingCard, ok := card.(*PlayingCard)
	if !ok {
		fmt.Printf("card received wasn't a playing card!")
		os.Exit(1)
	}

	fmt.Printf("card: %s\n", playingCard)
	fmt.Printf("  suit: %s\n", playingCard.Suit)
	fmt.Printf("  rank: %s\n", playingCard.Rank)
}
