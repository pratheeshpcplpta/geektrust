package app

import (
	"geektrust/app/db"
	"geektrust/app/models"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	db.DBConn
}

func InitApp() *App {
	connObje := App{}

	// check env file exists or not
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		panic("Unable to find .env file. Should create a file with name .env")
	}
	godotenv.Load()
	//initiate models

	connObje.LoadConfig(map[string]string{
		"Host":     os.Getenv("Host"),
		"User":     os.Getenv("User"),
		"Pass":     os.Getenv("Pass"),
		"Domine":   os.Getenv("Domine"),
		"Database": os.Getenv("Database"),
		"Port":     os.Getenv("Port"),
	})
	connObje.DBConnection() // Initiate a db connection with app

	models.AutoMigrateModel(connObje.DB) // Migrate model changes

	return &connObje
}
