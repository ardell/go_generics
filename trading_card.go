package main

type TradingCard struct {
	CollectibleName string
}

func NewTradingCard(collectibleName string) *TradingCard {
	return &TradingCard{CollectibleName: collectibleName}
}

func (tc *TradingCard) String() string {
	return tc.CollectibleName
}
