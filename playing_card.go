package main

import (
	"fmt"
)

type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, rank string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: rank}
}

func (pc *PlayingCard) String() string {
	return fmt.Sprintf("%s of %s", pc.Rank, pc.Suit)
}

func (pc *PlayingCard) Name() string {
	return fmt.Sprintf("[PlayingCard] %s", pc.String())
}
