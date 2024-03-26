package Controllers

import "github.com/gofiber/fiber/v2"

type User struct {
}

func (user User) DoSecret(c *fiber.Ctx) error {
	return nil
}
func (user User) DoPublic(c *fiber.Ctx) error {
	return nil
}
func (user User) Follow(c *fiber.Ctx) error {
	return nil
}
