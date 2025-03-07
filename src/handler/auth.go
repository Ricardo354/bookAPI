package handler

import (
	"errors"
	"livroAPI/src/database"
	"livroAPI/src/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func Register(c *fiber.Ctx) error{

	db := database.DBConn
	
	usuario := new(models.Usuario)

	if err := c.BodyParser(usuario); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Dados mal formados"})
	}

	result := db.Where("usuario = ?", usuario.Usuario).First(&usuario)

	if result.Error == nil{
		return c.Status(500).JSON(fiber.Map{"message" : "User already exists"})
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound){
		return c.Status(500).JSON(fiber.Map{"message":"Internal server error ocurred"})
	}
	
	hashed_password, err := HashPassword(usuario.Senha)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message":"Internal server error ocurred"})
	}
	usuario.Senha = hashed_password	

	if err := db.Create(usuario); err != nil{
		return c.Status(500).JSON(fiber.Map{"message":"Internal server error ocurred"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"Usuario criado com sucesso!"})

	// else, hash password and add to db

}
