package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

var Suits = []string{"Diamonds", "Hearts", "Spades", "Clubs"}
var Ranks = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

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

func (pc1 *PlayingCard) Compare(pc2 *PlayingCard) int {
	indexSuit1 := slices.IndexFunc(Suits, func(s string) bool { return s == pc1.Suit })
	indexSuit2 := slices.IndexFunc(Suits, func(s string) bool { return s == pc2.Suit })

	indexRank1 := slices.IndexFunc(Ranks, func(s string) bool { return s == pc1.Rank })
	indexRank2 := slices.IndexFunc(Ranks, func(s string) bool { return s == pc2.Rank })

	if indexSuit1 > indexSuit2 {
		return 1
	} else if indexSuit1 < indexSuit2 {
		return -1
	} else if indexRank1 > indexRank2 {
		return 1
	} else if indexRank1 < indexRank2 {
		return -1
	} else {
		return 0
	}
}
