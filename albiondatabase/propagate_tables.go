package albiondatabase

import (
	"albion/albionAPI"
	"albion/global"
	"database/sql"
	"log"
	"strings"
)

func getMarkets() []City {
	markets := []City{}
	for _, location := range WorldFileToSlice() {
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

var httpsCityResponse map[string]string = map[string]string{
	"Caerleon Market":       "Caerleon",
	"Thetford Market":       "Thetford",
	"Lymhurst Market":       "Lymhurst",
	"Bridgewatch Market":    "Bridgewatch",
	"Martlock Market":       "Martlock",
	"Fort Sterling Market":  "Fort Sterling",
	"Caerleon Black Market": "Caerleon 2",
	"Brecilien Market":      "5003",
}

func propagateCitiesTable(conn *sql.DB) {
	log.Println("Propagating Cities table...")

	tx, err := conn.Begin()
	if global.HandleErr(err) {
		return
	}

	stmt, err := tx.Prepare(
		"INSERT INTO Cities(idx, name, https_response_name) VALUES(?, ?, ?)")
	if global.HandleErr(err) {
		return
	}
	defer stmt.Close()

	for _, location := range getMarkets() {
		_, err = stmt.Exec(location.ID, location.NAME, httpsCityResponse[location.NAME])
		if global.HandleErr(err) {
			return
		}
	}
	err = tx.Commit()
	global.HandleErr(err)
}

var BasicItemsDescription []string = []string{
	"equipment item",
	"raw material",
	"refined material",
	"food",
}

var BasicItemsUniqueNames []string = []string{
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
	for _, substr := range BasicItemsDescription {
		if strings.Contains(strings.ToLower(description), substr) {
			return true
		}
	}
	return false
}

func doesUniqueNameContainAny(unique_name string) bool {
	for _, substr := range BasicItemsUniqueNames {
		if strings.Contains(strings.ToLower(unique_name), substr) {
			return true
		}
	}
	return false
}

func countItem(item Item) bool {
	return doesDescriptionContainAny(getEnDescription(item)) || doesUniqueNameContainAny(item.UNIQUE_NAME)
}

func getItems() []Item {
	items := []Item{}

	for _, item := range ItemsFileToSlice() {
		if countItem(item) {
			items = append(items, item)
		}
	}
	return items
}

func propagateItemsTable(conn *sql.DB) {
	log.Println("Propagating Items table...")

	tx, err := conn.Begin()
	if global.HandleErr(err) {
		return
	}

	stmt, err := tx.Prepare(
		"INSERT INTO Items(idx, name, tier, enchantment, quality) VALUES(?, ?, ?, ?, ?)")
	if global.HandleErr(err) {
		return
	}
	defer stmt.Close()

	for _, item := range getItems() {
		item := ItemToReadable(item)
		for _, quality := range QUALITIES {
			_, err = stmt.Exec(item.ID, item.NAME, item.TIER, item.ENCHANTMENT, quality)
			if global.HandleErr(err) {
				return
			}
		}
	}
	err = tx.Commit()
	global.HandleErr(err)
}

var NotFoundDatetime string = "0001-01-01 00:00:00"

func IsFoundOnMarket(item albionAPI.ItemOnMarket) bool {
	return (item.SELL_PRICE_MIN_DATE != NotFoundDatetime &&
		item.SELL_PRICE_MAX_DATE != NotFoundDatetime &&
		item.BUY_PRICE_MIN_DATE != NotFoundDatetime &&
		item.BUY_PRICE_MAX_DATE != NotFoundDatetime)
}

func GetItemDatabaseId(conn *sql.DB, item_idx string, quality int) int {
	rows, err := conn.Query(`
		SELECT id FROM Items WHERE
		idx = ? AND
		quality = ?
		`, item_idx, quality)
	if global.HandleErr(err) {
		return -1
	}
	defer rows.Close()

	if rows.Next() {
		var item_id int
		err = rows.Scan(&item_id)
		if global.HandleErr(err) {
			return -1
		}

		return item_id
	}

	return -1
}

func GetCityDatabaseId(conn *sql.DB, city_idx string) int {
	rows, err := conn.Query(`
		SELECT id FROM Cities WHERE
		https_response_name = ?
		`, city_idx)
	if global.HandleErr(err) {
		return -1
	}
	defer rows.Close()

	if rows.Next() {
		var city_id int
		err = rows.Scan(&city_id)
		if global.HandleErr(err) {
			return -1
		}

		return city_id
	}

	return -1
}

func propagateMarketsTable(conn *sql.DB) {
	log.Println("Propagating Markets table...")

	tx, err := conn.Begin()
	if global.HandleErr(err) {
		return
	}

	stmt, err := tx.Prepare(`
	INSERT INTO Markets(
		item_id,
		city_id,
		sell_price_min,
		sell_price_min_date,
		sell_price_max,
		sell_price_max_date,
		buy_price_min,
		buy_price_min_date,
		buy_price_max,
		buy_price_max_date
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if global.HandleErr(err) {
		return
	}
	defer stmt.Close()

	for _, item_on_market := range albionAPI.QueryAllItems() {
		if !IsFoundOnMarket(item_on_market) {
			continue
		}

		item_id := GetItemDatabaseId(
			conn, item_on_market.ITEM_IDX, item_on_market.QUALITY)
		city_id := GetCityDatabaseId(conn, item_on_market.CITY_IDX)
		_, err = stmt.Exec(
			item_id,
			city_id,
			item_on_market.SELL_PRICE_MIN,
			item_on_market.SELL_PRICE_MIN_DATE,
			item_on_market.SELL_PRICE_MAX,
			item_on_market.SELL_PRICE_MAX_DATE,
			item_on_market.BUY_PRICE_MIN,
			item_on_market.BUY_PRICE_MIN_DATE,
			item_on_market.BUY_PRICE_MAX,
			item_on_market.BUY_PRICE_MAX_DATE)
		if global.HandleErr(err) {
			return
		}
	}

	err = tx.Commit()
	global.HandleErr(err)
}

// ? Propagates Tables with common items
func PropagateTables() {
	log.Println("Propagating tables...")

	conn, err := InitializeDbConnection()
	if global.HandleErr(err) {
		return
	}

	propagateCitiesTable(conn)
	propagateItemsTable(conn)
	propagateMarketsTable(conn)

	log.Println("Tables propagated successfully")
}
