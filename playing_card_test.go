package main

import "testing"

type playingCardStringTest struct {
	suit, rank, expected string
}

var playingCardStringTests = []playingCardStringTest{
	playingCardStringTest{suit: "Spades", rank: "7", expected: "7 of Spades"},
	playingCardStringTest{suit: "Hearts", rank: "A", expected: "A of Hearts"},
	playingCardStringTest{suit: "Diamonds", rank: "K", expected: "K of Diamonds"},
	playingCardStringTest{suit: "Clubs", rank: "9", expected: "9 of Clubs"},
}

func TestPlayingCardStringReturnsExpectedValues(t *testing.T) {
	for _, test := range playingCardStringTests {
		playingCard := NewPlayingCard(test.suit, test.rank)
		if actual := playingCard.String(); actual != test.expected {
			t.Errorf("Expected %q, got %q", test.expected, actual)
		}
	}
}

type playingCardCompareTest struct {
	suit1, rank1, suit2, rank2 string
  expected int
}

var playingCardCompareTests = []playingCardCompareTest{
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "3",
    suit2: "Spades",
    rank2: "3",
    expected: 0,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "3",
    suit2: "Spades",
    rank2: "4",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "4",
    suit2: "Spades",
    rank2: "3",
    expected: 1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "A",
    suit2: "Spades",
    rank2: "2",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "9",
    suit2: "Spades",
    rank2: "J",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "J",
    suit2: "Spades",
    rank2: "Q",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "Q",
    suit2: "Spades",
    rank2: "K",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "A",
    suit2: "Spades",
    rank2: "K",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Diamonds",
    rank1: "3",
    suit2: "Hearts",
    rank2: "3",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Hearts",
    rank1: "3",
    suit2: "Spades",
    rank2: "3",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Spades",
    rank1: "3",
    suit2: "Clubs",
    rank2: "3",
    expected: -1,
  },
	playingCardCompareTest{
    suit1: "Diamonds",
    rank1: "K",
    suit2: "Hearts",
    rank2: "A",
    expected: -1,
  },
}

func TestPlayingCardCompareReturns0WhenCardsAreEqual(t *testing.T) {
	for _, test := range playingCardCompareTests {
		pc1 := NewPlayingCard(test.suit1, test.rank1)
		pc2 := NewPlayingCard(test.suit2, test.rank2)
		if actual := pc1.Compare(pc2); actual != test.expected {
			t.Errorf("When comparing %q with %q, expected %q, got %q", pc1, pc2, test.expected, actual)
		}
	}
}
