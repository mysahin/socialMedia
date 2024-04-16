package Controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"socialMedia/Database"
	"socialMedia/Helpers"
	"socialMedia/Models"
)

type User struct {
}

func (user User) DoSecret(c *fiber.Ctx) error {
	db := Database.DB.Db
	if Helpers.IsLogin(c) {
		cookie := c.Cookies("jwt")
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			c.Status(fiber.StatusUnauthorized)

			return c.JSON(fiber.Map{
				"error": "Giriş yapınız.",
			})
		}
		claims := token.Claims.(*jwt.StandardClaims)
		if err := db.Where("id=?", claims.Issuer).Update("is_secret", true); err != nil {
			return c.JSON(fiber.Map{
				"error": err,
			})
		}
	}

	return nil
}
func (user User) DoPublic(c *fiber.Ctx) error {
	db := Database.DB.Db
	if Helpers.IsLogin(c) {
		cookie := c.Cookies("jwt")
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			c.Status(fiber.StatusUnauthorized)

			return c.JSON(fiber.Map{
				"error": "Giriş yapınız.",
			})
		}
		claims := token.Claims.(*jwt.StandardClaims)
		if err := db.Where("id=?", claims.Issuer).Update("is_secret", false); err != nil {
			return c.JSON(fiber.Map{
				"error": err,
			})
		}
	}
	return c.JSON(fiber.Map{
		"error": "Lütfen giriş yapınız",
	})
}
func (user User) Follow(c *fiber.Ctx) error {
	db := Database.DB.Db
	var followed Models.User
	var follower Models.User
	if Helpers.IsLogin(c) {

		claims := getID(c)
		if err := db.First(&follower, "id=?", c.Params("id")); err != nil {
			return c.JSON(fiber.Map{
				"error": err,
			})
		}
		if err := db.First(&followed, "id=?", claims); err != nil {
			return c.JSON(fiber.Map{
				"error": err,
			})
		}
		follow := Models.Follow{
			FollowedUserName: followed.UserName,
			FollowerUserName: follower.UserName,
			IsFollow:         true,
		}
		if err := db.Create(&follow); err != nil {
			return c.JSON(fiber.Map{
				"error": err,
			})
		}

	}
	return c.JSON(fiber.Map{
		"error": "giriş yapınız.",
	})
}

func getID(c *fiber.Ctx) string {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)

		return ""
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer
}
