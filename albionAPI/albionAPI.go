package albionAPI

func GetAllMarketPrices() []ItemOnMarket {
	items_on_market := []ItemOnMarket{}

	for _, item_query := range prepareMarketPricesQueries() {
		query := AlbionApiURL + PricesPrefix + item_query + queryAll
		items_on_market = append(items_on_market, execMarketPricesQuery(query)...)
	}

	for _, item_on_market := range items_on_market {
		FormatDatetimes(&item_on_market)
	}

	return items_on_market
}
