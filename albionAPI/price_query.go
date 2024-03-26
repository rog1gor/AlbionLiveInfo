package albionAPI

import (
	"albion/global"
	"encoding/json"
	"net/http"
)

func prepareMarketPricesQueries() []string {
	item_queries := []string{}
	for _, item := range getItems() {
		if len(item_queries) == 0 {
			item_queries = append(item_queries, item.UNIQUE_NAME)
			continue
		}

		item_query_len :=
			len(item_queries[len(item_queries)-1]) + 1 + len(item.UNIQUE_NAME) //? query + , + item
		if item_query_len < AllTypesAndMarketsQuerySpaceLeft {
			item_queries[len(item_queries)-1] = item_queries[len(item_queries)-1] + "," + item.UNIQUE_NAME
		} else {
			item_queries = append(item_queries, item.UNIQUE_NAME)
		}
	}

	return item_queries
}

func execMarketPricesQuery(https_query string) []ItemOnMarket {
	items_on_market := []ItemOnMarket{}

	response, err := http.Get(https_query)
	if global.HandleErr(err) {
		return nil
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&items_on_market)
	if global.HandleErr(err) {
		return nil
	}

	return items_on_market
}

var CityResponse map[string]string = map[string]string{
	"Caerleon Market":       "Caerleon",
	"Thetford Market":       "Thetford",
	"Lymhurst Market":       "Lymhurst",
	"Bridgewatch Market":    "Bridgewatch",
	"Martlock Market":       "Martlock",
	"Fort Sterling Market":  "Fort Sterling",
	"Caerleon Black Market": "Caerleon 2",
	"Brecilien Market":      "5003",
}
