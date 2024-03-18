package albionAPI

import (
	"albion/global"
	"encoding/json"
	"net/http"
)

func getAllItemIdxs() []string {
	conn, err := InitializeDbConnection()
	if global.HandleErr(err) {
		return nil
	}

	rows, err := conn.Query("SELECT DISTINCT idx FROM Items;")
	if global.HandleErr(err) {
		return nil
	}
	defer rows.Close()

	item_idxs := []string{}
	for rows.Next() {
		var idx string
		err = rows.Scan(&idx)
		if global.HandleErr(err) {
			return nil
		}
		item_idxs = append(item_idxs, idx)
	}

	return item_idxs
}

// ? Takes all item idxs and forms queries so that they
// ? don't exceed https request character limit
func prepareItemQueries(item_idxs []string) []string {
	item_queries := []string{}
	for _, item := range item_idxs {
		if len(item_queries) == 0 {
			item_queries = append(item_queries, item)
			continue
		}

		item_query_len := len(item_queries[len(item_queries)-1]) + 1 + len(item) //? query + , + item
		if item_query_len < AllTypesAndMarketsQuerySpaceLeft {
			item_queries[len(item_queries)-1] = item_queries[len(item_queries)-1] + "," + item
		} else {
			item_queries = append(item_queries, item)
		}
	}

	return item_queries
}

func ExecQuery(https_query string) []ItemOnMarket {
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

func QueryAllItems() []ItemOnMarket {
	items_on_market := []ItemOnMarket{}

	item_idxs := getAllItemIdxs()
	item_queries := prepareItemQueries(item_idxs)

	for _, item_query := range item_queries {
		query := AlbionApiURL + PricesPrefix + item_query + queryAll
		items_on_market = append(items_on_market, ExecQuery(query)...)
	}

	for _, item_on_market := range items_on_market {
		FormatDatetimes(&item_on_market)
	}

	return items_on_market
}
