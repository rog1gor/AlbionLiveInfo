package albiondatabase

import (
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

func propagateCitiesTable(conn *sql.DB) {
	log.Println("Propagating Cities table...")

	tx, err := conn.Begin()
	if global.HandleErr(err) {
		return
	}

	stmt, err := tx.Prepare("INSERT INTO Cities(IDX, NAME) VALUES(?, ?)")
	if global.HandleErr(err) {
		return
	}
	defer stmt.Close()

	for _, location := range getMarkets() {
		_, err = stmt.Exec(location.ID, location.NAME)
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
		"INSERT INTO Items(IDX, NAME, TIER, ENCHANTMENT) VALUES(?, ?, ?, ?)")
	if global.HandleErr(err) {
		return
	}
	defer stmt.Close()

	for _, item := range getItems() {
		item := ItemToReadable(item)
		_, err = stmt.Exec(item.ID, item.NAME, item.TIER, item.ENCHANTMENT)
		if global.HandleErr(err) {
			return
		}
	}
	err = tx.Commit()
	global.HandleErr(err)
}

func propagateMarketsTable(conn *sql.DB) {
	log.Println("Propagating Marketsc table...")
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
