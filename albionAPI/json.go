package albionAPI

import (
	"albion/global"
	"encoding/json"
	"os"
	"strings"
)

var WORLD_JSON string = pkgPath + "JSON/world.json"
var ITEMS_JSON string = pkgPath + "JSON/items.json"

// & Markets

func worldFileToSlice() []City {
	world_file, err := os.Open(WORLD_JSON)
	if global.HandleErr(err) {
		return nil
	}
	defer world_file.Close()

	var cities []City
	err = json.NewDecoder(world_file).Decode(&cities)
	if global.HandleErr(err) {
		return nil
	}
	return cities
}

func getMarkets() []City {
	markets := []City{}
	for _, location := range worldFileToSlice() {
		if strings.Contains(strings.ToLower(location.NAME), "market") {
			// ? Black Market corner case
			if strings.Contains(strings.ToLower(location.ID), "auction") {
				location.NAME = "Caerleon Black Market"
			}
			markets = append(markets, location)
		}
	}
	return markets
}

// & Items

func itemsFileToSlice() []Item {
	items_file, err := os.Open(ITEMS_JSON)
	if global.HandleErr(err) {
		return nil
	}
	defer items_file.Close()

	var items []Item
	err = json.NewDecoder(items_file).Decode(&items)
	if global.HandleErr(err) {
		return nil
	}
	return items
}

var basicItemsDescriptionFilter []string = []string{
	"equipment item",
	"raw material",
	"refined material",
	"food",
}

var basicItemsUniqueNamesFilter []string = []string{
	"_mount",
	"_faction",
	"_baby",
	"_fish",
	"_rune",
	"_relic",
	"_soul",
	"_artefact",
}

func doesDescriptionContainAny(description string) bool {
	for _, substr := range basicItemsDescriptionFilter {
		if strings.Contains(strings.ToLower(description), substr) {
			return true
		}
	}
	return false
}

func doesUniqueNameContainAny(unique_name string) bool {
	for _, substr := range basicItemsUniqueNamesFilter {
		if strings.Contains(strings.ToLower(unique_name), substr) {
			return true
		}
	}
	return false
}

func countItem(item Item) bool {
	return doesDescriptionContainAny(item.getEnDescription()) || doesUniqueNameContainAny(item.UNIQUE_NAME)
}

func getItems() []Item {
	items := []Item{}

	for _, item := range itemsFileToSlice() {
		if countItem(item) {
			items = append(items, item)
		}
	}
	return items
}
