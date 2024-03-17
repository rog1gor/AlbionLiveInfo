package albiondatabase

import (
	"albion/global"
	"log"
	"os"
)

func RemoveDatabase() {
	err := os.Remove(AlbionDatabase)
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	log.Println("Database removed successfully.")
}

func CreateDatabase() {
	if _, err := os.Stat(AlbionDatabase); err == nil {
		log.Println("Database already exists.")
	} else if os.IsNotExist(err) {
		db, err := InitializeDbConnection()
		global.HandleErr(err)
		defer db.Close()

		log.Println("Database created successfully.")
	} else {
		log.Fatalln(err)
	}
}
