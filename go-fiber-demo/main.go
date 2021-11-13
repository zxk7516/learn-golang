package main

import (
	"fmt"
	"runtime"
	"time"
	"zxk/go-fiber-demo/common"
	"zxk/go-fiber-demo/controller/admin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
func private(ctx *fiber.Ctx) (err error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)
	ctx.Send([]byte(fmt.Sprintf("hello user %s", id)))
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

	err := common.LoadConfig("config.yml")
	if err != nil {
		panic(err.Error())
	}
	err = common.InitDb()
	if err != nil {
		panic(err.Error())
	}
	common.TestPgDb()
	// models.MakeMigrations()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost, http://localhost:3000, http://127.0.0.1, https://gofiber.net",
	}))
	app.Get("/", hello)
	app.Post("/login", login)
	app.Post("/register", login)

	jwtWare := jwtware.New(jwtware.Config{
		SigningKey: []byte(jwt_secret),
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "Unauthorized",
				"detail": e.Error(),
			})
			return nil
		},
	})

	app.Get("/private", jwtWare, private)
	app.Get("/nonPrivate", hello)

	adminJwt := admin.GetAdminJwtWare()
	adminRoute := app.Group("/admin", adminJwt)
	adminRoute.Post("/login", admin.AdminLogin)
	adminRoute.Get("/users", admin.AdminUsers)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
