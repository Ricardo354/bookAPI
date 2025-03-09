package handler

import (
	"livroAPI/src/database"
	"livroAPI/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllLivros(c *fiber.Ctx) error {

	db := database.DBConn

	var Livros []models.Livro

	if err := db.Find(&Livros).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar Livros",
		})
	}

	return c.Status(fiber.StatusOK).JSON(Livros)
}

func GetLivroByID(c *fiber.Ctx) error {

	db := database.DBConn

	Livro := new(models.Livro)

	result := db.Find(Livro, c.Params("id"))
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar Livro",
		})
	} else if result.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(Livro)
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Livro não encontrado!",
		})
	}

}

func CreateLivro(c *fiber.Ctx) error {

	db := database.DBConn

	Livro := new(models.Livro)

	if err := c.BodyParser(Livro); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao interpretar dados",
		})
	}

	// validator

	if err := db.Create(Livro).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao criar Livro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Livro criado": Livro,
	})

}

func UpdateLivroByID(c *fiber.Ctx) error {

	db := database.DBConn

	Livro := new(models.Livro)

	if err := db.First(Livro, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Livro não encontrado",
		})
	}

	if err := c.BodyParser(Livro); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao interpretar dados",
		})
	}

	if err := db.Save(Livro).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao atualizar livro",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Livro atualizado com sucesso",
			"Livro":   Livro})
	}

}
func DeleteLivroByID(c *fiber.Ctx) error {

	db := database.DBConn

	id := c.Params("id")

	if err := db.Delete(&models.Livro{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Erro": "Erro ao deletar livro",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ID Deletado": id,
	})

}
