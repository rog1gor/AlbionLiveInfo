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
