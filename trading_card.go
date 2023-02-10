package main

import (
	"fmt"
)

type TradingCard struct {
	CollectibleName string
}

func NewTradingCard(collectibleName string) *TradingCard {
	return &TradingCard{CollectibleName: collectibleName}
}

func (tc *TradingCard) String() string {
	return tc.CollectibleName
}

func (tc *TradingCard) Name() string {
	return fmt.Sprintf("[TradingCard] %s", tc.String())
}
