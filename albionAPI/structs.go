package albionAPI

import "strings"

type ItemOnMarket struct {
	ITEM_IDX string `json:"item_id"`
	CITY_IDX string `json:"city"`
	QUALITY  int    `json:"quality"`

	SELL_PRICE_MIN      int    `json:"sell_price_min"`
	SELL_PRICE_MIN_DATE string `json:"sell_price_min_date"`

	SELL_PRICE_MAX      int    `json:"sell_price_max"`
	SELL_PRICE_MAX_DATE string `json:"sell_price_max_date"`

	BUY_PRICE_MIN      int    `json:"buy_price_min"`
	BUY_PRICE_MIN_DATE string `json:"buy_price_min_date"`

	BUY_PRICE_MAX      int    `json:"buy_price_max"`
	BUY_PRICE_MAX_DATE string `json:"buy_price_max_date"`
}

func FormatDatetimes(item *ItemOnMarket) {
	item.SELL_PRICE_MIN_DATE = strings.ReplaceAll(
		item.SELL_PRICE_MIN_DATE, "T", " ")
	item.SELL_PRICE_MAX_DATE = strings.ReplaceAll(
		item.SELL_PRICE_MAX_DATE, "T", " ")
	item.BUY_PRICE_MIN_DATE = strings.ReplaceAll(
		item.BUY_PRICE_MIN_DATE, "T", " ")
	item.BUY_PRICE_MAX_DATE = strings.ReplaceAll(
		item.BUY_PRICE_MAX_DATE, "T", " ")
}
