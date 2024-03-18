package albiondatabase

import (
	"albion/global"
	"database/sql"
	"log"
)

func createCitiesTable(conn *sql.DB) {
	log.Println("Creating table for Cities...")

	sqlStmt := `
	DROP TABLE IF EXISTS Cities;
	CREATE TABLE Cities (
		id integer not null primary key,
		idx text not null,
		name text not null,
		https_response_name text not null
	);
	DELETE FROM Cities;`

	_, err := conn.Exec(sqlStmt)
	if global.HandleErr(err) {
		return
	}
}

func createItemsTable(conn *sql.DB) {
	log.Println("Creating table for Items...")

	sqlStmt := `
	DROP TABLE IF EXISTS Items;
	CREATE TABLE Items (
		id integer not null primary key,
		idx text not null,
		name text not null,
		tier integer not null,
		enchantment integer,
		quality integer not null
	);
	DELETE FROM Items;`
	_, err := conn.Exec(sqlStmt)
	if global.HandleErr(err) {
		return
	}
}

func createMarketTable(conn *sql.DB) {
	log.Println("Creating table for Markets...")

	sqlStmt := `
	DROP TABLE IF EXISTS Markets;
	CREATE TABLE Markets (
		item_id integer,
		city_id integer,
		sell_price_min integer,
		sell_price_min_date datetime,
		sell_price_max integer,
		sell_price_max_date datetime,
		buy_price_min integer,
		buy_price_min_date datetime,
		buy_price_max integer,
		buy_price_max_date datetime,
		PRIMARY KEY (item_id, city_id),
		FOREIGN KEY (item_id) REFERENCES Items(id),
		FOREIGN KEY (city_id) REFERENCES Cities(id)
	);
	DELETE FROM Markets`
	_, err := conn.Exec(sqlStmt)
	if global.HandleErr(err) {
		return
	}
}

func CreateTables() {
	log.Println("Creating tables...")

	conn, err := InitializeDbConnection()
	if global.HandleErr(err) {
		return
	}

	createCitiesTable(conn)
	createItemsTable(conn)
	createMarketTable(conn)

	log.Println("Tables created successfully")
}
