package marketevaluation

import (
	"albion/albionAPI"
	"math"
)

func findBestResellForItem(item_on_markets []albionAPI.ItemOnMarket) Resell {
	buy_city, buy_price := "None", math.MaxInt
	sell_city, sell_price := "None", 0

	for _, item := range item_on_markets {
		if buy_price > item.SELL_PRICE_MIN {
			buy_price = item.SELL_PRICE_MIN
			buy_city = item.CITY_IDX
		}

		if sell_price < item.BUY_PRICE_MAX {
			sell_price = item.BUY_PRICE_MAX
			sell_city = item.CITY_IDX
		}
	}

	return Resell{
		BUY_CITY:   buy_city,
		BUY_PRICE:  buy_price,
		SELL_CITY:  sell_city,
		SELL_PRICE: sell_price,
	}
}
