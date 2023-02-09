package main

import "fmt"

func main() {
	deck := NewPlayingCardDeck()

	playingCard := deck.RandomCard()
	fmt.Printf("card: %s\n", playingCard)
	fmt.Printf("  suit: %s\n", playingCard.Suit)
	fmt.Printf("  rank: %s\n", playingCard.Rank)
}
