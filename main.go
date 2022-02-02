package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-todo/database"
	"go-todo/todo"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	database.ConnectDB()

	sqlDB, _ := database.DB.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			println("Hit error when obtaining reference to sqlDB: " + err.Error())
		}
	}(sqlDB)

	api := app.Group("/api")
	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":5001"))
}
