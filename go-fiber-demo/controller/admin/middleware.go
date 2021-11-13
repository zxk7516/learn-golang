package admin

import (
	"zxk/go-fiber-demo/common"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

func GetAdminJwtWare() fiber.Handler {
	adminJwtWare :=
		jwt.New(jwt.Config{
			Filter: func(c *fiber.Ctx) bool {
				return c.Path() == "/admin/login"
			},
			ContextKey: "admin",
			SigningKey: []byte(common.GetConfig().AdminJwtSecret),
			SuccessHandler: func(c *fiber.Ctx) error {
				c.Next()
				return nil
			},
			ErrorHandler: func(c *fiber.Ctx, e error) error {
				c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":  "Unauthorized",
					"detail": e.Error(),
				})
				return nil
			},
		})
	return adminJwtWare
}
