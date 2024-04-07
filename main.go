package main

import (
	"albion/albionAPI"
	"albion/marketevaluation"
	"fmt"
)

func main() {
	items_on_markets := albionAPI.GetAllMarketPrices()
	resell_item, resell := marketevaluation.FindBestResell(items_on_markets)
	fmt.Println("Best resell found!")
	fmt.Println("Item:", resell_item)
	fmt.Printf("Where to buy: %s (for %d)\n", resell.BUY_CITY, resell.BUY_PRICE)
	fmt.Printf("Where to sell: %s (for %d)\n", resell.SELL_CITY, resell.SELL_PRICE)
}
