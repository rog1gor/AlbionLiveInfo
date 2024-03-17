package albiondatabase

func ResetDatabase() {
	RemoveDatabase()
	CreateDatabase()
	CreateTables()
	PropagateTables()
}
