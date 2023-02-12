package main

import "testing"

type tradingCardStringTest struct {
	collectibleName, expected string
}

var stringTests = []tradingCardStringTest{
	tradingCardStringTest{collectibleName: "Foo", expected: "Foo"},
	tradingCardStringTest{collectibleName: "Bar", expected: "Bar"},
	tradingCardStringTest{collectibleName: "Baz", expected: "Baz"},
}

func TestStringReturnsExpectedValues(t *testing.T) {
	for _, test := range stringTests {
		card := NewTradingCard(test.collectibleName)
		if actual := card.String(); actual != test.expected {
			t.Errorf("Expected %q, got %q", test.expected, actual)
		}
	}
}
