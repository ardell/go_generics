package main

import "fmt"

func main() {
	playingCardDeck := NewPlayingCardDeck()
	playingCard := playingCardDeck.RandomCard()
	fmt.Printf("playing card: %s\n", playingCard)
	fmt.Printf("  suit: %s\n", playingCard.Suit)
	fmt.Printf("  rank: %s\n", playingCard.Rank)

	tradingCardDeck := NewTradingCardDeck()
	tradingCard := tradingCardDeck.RandomCard()
	fmt.Printf("trading card: %s\n", tradingCard)
	fmt.Printf("  collectible name: %s\n", tradingCard.CollectibleName)
}
