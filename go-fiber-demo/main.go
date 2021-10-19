package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

const jwt_secret string = "jwt_secret"

func authRequired() func(ctx *fiber.Ctx) error {

	return nil

}

func hello(ctx *fiber.Ctx) (err error) {
	ctx.Send([]byte("hello world"))
	return nil
}

func login(ctx *fiber.Ctx) (err error) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}
	if body.Email != "bob@gmail.com" || body.Password != "password" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	s, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		fmt.Println(err)
		ctx.SendStatus(fiber.StatusInternalServerError)
		return
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}{
			Id:    1,
			Email: "bob@gmail.com",
		},
	})
	return
}
func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", hello)
	app.Post("/login", login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(jwt_secret),
	}))

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
