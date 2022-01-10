package controllers

import (
	"APICHATGOLANG/database"
	"APICHATGOLANG/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetAllMessages(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"messages": "belum login, silahkan login",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var pesan models.Messages

	database.DB.Find(&pesan, claims.Issuer)

	return c.JSON(&pesan)
}

func MakeMessages(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"messages": "belum login, silahkan login",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var pesan map[string]string

	if err := c.BodyParser(&pesan); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	nopengirim := claims.Issuer

	pesanUI := models.Messages{
		No_Pengirim: nopengirim,
		Pesan:       pesan["pesan"],
		No_Penerima: pesan["penerima"],
	}

	database.DB.Create(&pesanUI)
	return c.JSON(&pesanUI)
}
