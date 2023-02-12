package main

import "testing"

func TestPlayingCardDeckContains52Cards(t *testing.T) {
	deck := NewPlayingCardDeck()
	const expected = 52
	actual := len(deck.cards)
	if actual != expected {
		t.Errorf("Expected playing card deck to have %q cards, go %q", expected, actual)
	}
}

func TestNewTradingCardDeckCanBeConstructed(t *testing.T) {
	// Not a super useful test, but we didn't want to test the number of cards,
	// since it's likely we'll add more `collectibles` in the future.
	NewTradingCardDeck()
}
