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
		ID integer not null primary key,
		IDX text not null,
		NAME text not null
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
		ID integer not null primary key,
		IDX text not null,
		NAME text not null,
		TIER integer,
		ENCHANTMENT integer
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
		quantity integer,
		current_cost integer,
		min_cost_7days integer,
		max_cost_7days integer,
		avg_cost_7days integer,
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
