package handler

import "tugas-sesi12/database"

func StartApp() {
	var port = "8080"
	database.InitializeDatabase()

	db := database.GetDatabaseInstance()

	_, _ = port, db

}
