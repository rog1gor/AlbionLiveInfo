package albiondatabase

import "albion/albionAPI"

func ResetDatabase() {
	RemoveDatabase()
	CreateDatabase()
	CreateTables()
	PropagateTables()
	albionAPI.QueryAllItems()
}
