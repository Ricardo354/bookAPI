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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckUsernameExists(u string) (bool, error) {
    db := database.DBConn
    usuario := new(models.Usuario)

    if err := db.Where("usuario = ?", u).First(usuario).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, nil
        }
        return false, err
    }
    return true, nil 
}

func Register(c *fiber.Ctx) error {

	db := database.DBConn

	usuario := new(models.Usuario)

	if err := c.BodyParser(usuario); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dados mal formados"})
	}

	exists, err := CheckUsernameExists(usuario.Username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erro ao verificar usuário"})
	}

	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Usuário já existe"})
	}

	hashed_password, err := HashPassword(usuario.Senha)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error ocurred"})
	}
	usuario.Senha = hashed_password

	if err := db.Create(usuario).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error ocurred"})
	}

	return c.Status(fiber.StatusCreated).JSON(usuario)


}

