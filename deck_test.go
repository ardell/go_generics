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

func TestAddCardAddsACard(t *testing.T) {
	deck := NewPlayingCardDeck()
	newCard := NewPlayingCard("spades", "9")

	before := len(deck.cards)
	deck.AddCard(newCard)
	after := len(deck.cards)

	actualChange := after - before
	const expectedChange = 1
	if actualChange != 1 {
		t.Errorf("Expected number of cards in deck to change by %q, actually changed by %q", expectedChange, actualChange)
	}
}

func TestIndexOfCardReturnsIndexOfAGivenCardWhenCardExists(t *testing.T) {
	deck := &Deck[*PlayingCard]{}
	cardToSearchFor := NewPlayingCard("spades", "3")
	deck.AddCard(NewPlayingCard("spades", "9"))
	deck.AddCard(NewPlayingCard("hearts", "3"))
	deck.AddCard(cardToSearchFor)
	deck.AddCard(NewPlayingCard("diamonds", "J"))

	actualIndex, ok := deck.IndexOfCard(cardToSearchFor)
	const expectedIndex = 2

	if ok != nil {
		t.Errorf("Expected IndexOfCard to return ok, but it returned error.")
	}

	if actualIndex < 0 {
		t.Errorf("Expected card to be at index %q, but got %q", expectedIndex, actualIndex)
	}
}

func TestIndexOfCardReturnsErrorWhenCardDoesNotExist(t *testing.T) {
	deck := &Deck[*PlayingCard]{}
	cardToSearchFor := NewPlayingCard("spades", "3")
	deck.AddCard(NewPlayingCard("spades", "9"))
	deck.AddCard(NewPlayingCard("hearts", "3"))
	deck.AddCard(NewPlayingCard("diamonds", "J"))

	_, ok := deck.IndexOfCard(cardToSearchFor)
	if ok == nil {
		t.Errorf("Expected IndexOfCard to return error, but it returned ok.")
	}
}
