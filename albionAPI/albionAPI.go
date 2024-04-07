package albionAPI

import "fmt"

func GetAllMarketPrices() []ItemOnMarket {
	items_on_markets := []ItemOnMarket{}

	fmt.Println("Querying albionAPI...")
	item_queries := prepareMarketPricesQueries()
	for i, item_query := range item_queries {
		query := AlbionApiURL + PricesPrefix + item_query + queryAll
		items_on_markets = append(items_on_markets, execMarketPricesQuery(query)...)
		fmt.Printf("\033[1A\033[K")
		fmt.Printf("Query rogress %d out of %d...\n", i+1, len(item_queries))
	}

	for _, item_on_market := range items_on_markets {
		FormatDatetimes(&item_on_market)
	}

	return items_on_markets
}
