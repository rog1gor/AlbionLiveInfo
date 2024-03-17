package albiondatabase

import (
	"albion/global"
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var ALBIONDATABASE_PATH string = "albiondatabase/"

var AlbionDatabase string = "Albion.db"

var WORLD_JSON string = ALBIONDATABASE_PATH + "JSON/world.json"
var ITEMS_JSON string = ALBIONDATABASE_PATH + "JSON/items.json"

var QUALITIES map[string]int = map[string]int{
	"normal":      1,
	"good":        2,
	"outstanding": 3,
	"excellent":   4,
	"masterpiece": 5,
}

var CITY_MAPPING map[string]string = map[string]string{
	"5003":          "5003",
	"Black Market":  "3013-Auction2",
	"Bridgewatch":   "2004",
	"Caerleon":      "3005",
	"Fort Sterling": "4002",
	"Lymhurst":      "1002",
	"Martlock":      "3008",
	"Thetford":      "0007",
}

func InitializeDbConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", AlbionDatabase)
}

func WorldFileToSlice() []City {
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

func ItemsFileToSlice() []Item {
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
