package routes

import (
	"livroAPI/src/handler"

	"github.com/gofiber/fiber/v2"
)

//JWT auth
//Rate limiting
//godotenv
//docker
// unit testing
// fiber.requestcode

func LivroRoutes(app *fiber.App) {

	book_endpoints := app.Group("/livro")

	book_endpoints.Get("/", handler.GetAllLivros)
	book_endpoints.Get("/:id", handler.GetLivroByID)
	book_endpoints.Post("/", handler.CreateLivro)
	book_endpoints.Put("/:id", handler.UpdateLivroByID)
	book_endpoints.Delete("/:id", handler.DeleteLivroByID)

	auth_endpoints := app.Group("/auth")

	auth_endpoints.Post("/signup", handler.Register)
	auth_endpoints.Post("/login", handler.Login)

}
