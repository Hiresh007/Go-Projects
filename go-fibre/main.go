package main

import (
	"fmt"
	"go-fibre/database"
	"go-fibre/leads"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", leads.GetLeads)
	app.Get("/api/v1/lead/:id", leads.GetLead)
	app.Post("/api/v1/lead", leads.NewLead)
	app.Delete("/api/v1/lead/:id", leads.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBCon, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed  to connect to DB")

	}

	fmt.Println("Connection opened to database")
	database.DBCon.AutoMigrate(&leads.Lead{})

}
func main() {
	app := fiber.New()
	initDatabase()
	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
	defer database.DBCon.Close()
}
