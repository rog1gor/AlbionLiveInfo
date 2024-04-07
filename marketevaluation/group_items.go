package marketevaluation

import (
	"albion/albionAPI"
	"strconv"
)

func mapItems(items_on_markets []albionAPI.ItemOnMarket) map[string][]albionAPI.ItemOnMarket {
	item_map := make(map[string][]albionAPI.ItemOnMarket)
	for _, item := range items_on_markets {
		key := item.ITEM_IDX + "#" + strconv.Itoa(item.QUALITY)
		item_map[key] = append(item_map[key], item)
	}
	return item_map
}
