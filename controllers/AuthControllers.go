package controllers

import (
	"APICHATGOLANG/database"
	"APICHATGOLANG/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var Secretkey = "secret"

func LoginAuth(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("no_handphone = ?", data["no_handphone"]).First(&user)

	if user.NoHandphone == "" {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"messages": "no belum terdaftar",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"messages": "salah password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.NoHandphone,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(Secretkey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"messages": "tidak bisa login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"messages": "sukses",
	})
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	passwordUI, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	userUI := models.User{
		Nama:        data["nama"],
		NoHandphone: data["no_handphone"],
		Password:    passwordUI,
	}

	database.DB.Create(&userUI)
	return c.JSON(&userUI)
}

func GetUser(c *fiber.Ctx) error {

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

	var user models.User

	database.DB.Where("no_handphone = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}
