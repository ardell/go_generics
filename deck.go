package main

import (
	"math/rand"
	"time"
)

type Deck[C Card] struct {
	cards []C
}

func NewPlayingCardDeck() *Deck[*PlayingCard] {
	suits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := &Deck[*PlayingCard]{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	return deck
}

func NewTradingCardDeck() *Deck[*TradingCard] {
	collectibles := []string{"Foo", "Bar", "Baz", "Fizz", "Buzz"}

	deck := &Deck[*TradingCard]{}
	for _, collectibleName := range collectibles {
		deck.AddCard(NewTradingCard(collectibleName))
	}
	return deck
}

func (d *Deck[C]) AddCard(card C) {
	d.cards = append(d.cards, card)
}

func (d *Deck[C]) RandomCard() C {
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	cardIndex := randomNumberGenerator.Intn(len(d.cards))
	return d.cards[cardIndex]
}
