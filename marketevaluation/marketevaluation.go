package marketevaluation

import (
	"albion/albionAPI"
	"math"
)

type Resell struct {
	BUY_CITY   string
	BUY_PRICE  int
	SELL_CITY  string
	SELL_PRICE int
}

func (r Resell) CalcIncome() int {
	return r.SELL_PRICE - r.BUY_PRICE
}

func FindBestResell(items_on_markets []albionAPI.ItemOnMarket) (string, Resell) {
	mapped_items := mapItems(items_on_markets)
	best_item_to_resell := "None"
	best_resell := Resell{
		BUY_CITY:   "None",
		BUY_PRICE:  math.MaxInt,
		SELL_CITY:  "None",
		SELL_PRICE: 0,
	}

	for item, item_on_markets := range mapped_items {
		resell := findBestResellForItem(item_on_markets)
		if resell.CalcIncome() > best_resell.CalcIncome() {
			best_resell = resell
			best_item_to_resell = item
		}
	}

	return best_item_to_resell, best_resell
}
