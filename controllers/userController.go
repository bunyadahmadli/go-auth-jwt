package controllers

import (
	"auth-go/database"
	"auth-go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "unauthenticated"})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.Db.Where("id = ?", claims.Issuer).First(&user)
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {

	return nil
}
