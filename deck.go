package main

import (
	"errors"
	"golang.org/x/exp/slices"
	"math/rand"
	"time"
)

type Deck[C Card] struct {
	cards []C
}

func NewPlayingCardDeck() *Deck[*PlayingCard] {
	deck := &Deck[*PlayingCard]{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
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

func (d *Deck[C]) IndexOfCard(card C) (int, error) {
	index := slices.IndexFunc(d.cards, func(c C) bool { return c == card })
	if index >= 0 {
		return index, nil
	} else {
		return 0, errors.New("Card not found")
	}
}

func (d *Deck[C]) RandomCard() C {
	randomNumberGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	cardIndex := randomNumberGenerator.Intn(len(d.cards))
	return d.cards[cardIndex]
}
