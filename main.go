package main

import "fmt"

func printCard[C Card](card C) {
	fmt.Printf("card: %s\n", card)
	fmt.Printf("  name: %s\n", card.Name())
}

func main() {
	playingCardDeck := NewPlayingCardDeck()
	playingCard := playingCardDeck.RandomCard()
	printCard(playingCard)
	fmt.Printf("  suit: %s\n", playingCard.Suit)
	fmt.Printf("  rank: %s\n", playingCard.Rank)

	tradingCardDeck := NewTradingCardDeck()
	tradingCard := tradingCardDeck.RandomCard()
	printCard(tradingCard)
	fmt.Printf("  collectible name: %s\n", tradingCard.CollectibleName)
}
