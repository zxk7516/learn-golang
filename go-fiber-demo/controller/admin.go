package controller

import (
	"fmt"
	"time"
	"zxk/go-fiber-demo/common"
	"zxk/go-fiber-demo/models/admin"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AdminLogin(ctx *fiber.Ctx) (err error) {

	type login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var l login
	err = ctx.BodyParser(&l)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}
	adminUser, err := admin.GetAdminUser(admin.AdminUsers{UserName: l.Username})
	if err != nil {
		fmt.Println(err)
	}
	valid := common.CheckPassword(l.Password, adminUser.Password)
	if !valid {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password valid failed",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = adminUser.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	s, err := token.SignedString([]byte(common.GetConfig().AdminJwtSecret))
	if err != nil {
		fmt.Println(err)
		ctx.SendStatus(fiber.StatusInternalServerError)
		return
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"admin": adminUser,
	})
	return
}
