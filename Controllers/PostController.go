package Controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"socialMedia/Database"
	"socialMedia/Helpers"
	"socialMedia/Models"
)

type Post struct {
}

func (post Post) Share(c *fiber.Ctx) error {
	db := Database.DB.Db
	_post := new(Models.Post)
	user := new(Models.User)
	if err := c.BodyParser(&_post); err != nil {
		return err
	}
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	db.First(&user, "id=?", claims.Issuer)
	newPost := Models.Post{
		UserName:  user.UserName,
		Write:     _post.Write,
		IsArchive: false,
	}

	if err := db.Create(&newPost).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Yeni gönderi başarıyla paylaşıldı.",
		"post":    newPost,
	})

	/*return c.JSON(fiber.Map{
		"message": "hata",
	})*/
}
func (post Post) Delete(c *fiber.Ctx) error {
	db := Database.DB.Db
	postID := c.Params("id")
	var dPost Models.Post
	if err := db.First(&dPost, postID).Error; err != nil {
		return err
	}
	if err := db.Delete(&dPost).Error; err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "Gönderi başarıyla silindi.",
	})
}
func (post Post) Archive(c *fiber.Ctx) error {
	db := Database.DB.Db
	postID := c.Params("id")
	var archivedPost Models.Post
	if err := db.First(&archivedPost, postID).Error; err != nil {
		return err
	}
	if err := db.Model(&archivedPost).Where("id=?", postID).Update("is_archive", "1").Error; err != nil {
		return err
	}
	return nil
}
func (post Post) UnArchive(c *fiber.Ctx) error {
	db := Database.DB.Db
	postID := c.Params("id")
	var archivedPost Models.Post
	if err := db.First(&archivedPost, postID).Error; err != nil {
		return err
	}
	if err := db.Model(&archivedPost).Where("id=?", postID).Update("is_archive", "0").Error; err != nil {
		return err
	}
	return nil
}

func (post Post) ViewPosts(c *fiber.Ctx) error {
	isLogin := Helpers.IsLogin(c)
	if isLogin {
		db := Database.DB.Db
		var viewPost []Models.Post
		var users []Models.User
		var followed []Models.Follow
		var follower []Models.Follow
		var user Models.User
		cookie := c.Cookies("jwt")
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err,
			})
		}
		claims := token.Claims.(*jwt.StandardClaims)
		db.First(&user, "id=?", claims.Issuer)
		db.Find(&users)
		db.Find(&follower)
		for _, y := range users {
			db.Find(&followed, "followed_user_name=? AND follower_user_name=?", y.UserName, user.UserName)
		}
		for _, b := range followed {
			if err := db.Find(&viewPost, "is_archive=? AND user_name=?", false, b.FollowedUserName).Error; err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{
			"Postlar": viewPost,
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Önce giriş yapmalısınız!!!",
	})
}
