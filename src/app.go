package main

import (
	"livroAPI/src/database"
	"livroAPI/src/models"
	"livroAPI/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	
	database.InitDatabase()
	database.DBConn.AutoMigrate(&models.Livro{})
	database.DBConn.AutoMigrate(&models.Usuario{})
	

	file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
		DisableColors: true,
	}))


	routes.LivroRoutes(app)


	app.Listen(":3000")

}
