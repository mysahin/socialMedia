package Controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"socialMedia/Database"
	"socialMedia/Helpers"
	"socialMedia/Models"
	"strconv"
)

type Post struct {
}

func (post Post) Share(c *fiber.Ctx) error {
	fileName, err := UploadFile(c)
	if err != nil {
		return c.JSON("hata var!!!")
	}
	fmt.Println(fileName)
	isLogin := Helpers.IsLogin(c)
	if isLogin {
		db := Database.DB.Db
		_post := new(Models.Post)
		user := new(Models.User)
		if err := c.BodyParser(&_post); err != nil {
			return err
		}
		id := getID(c)
		db.First(&user, "id=?", id)
		newPost := Models.Post{
			UserName:  user.UserName,
			Write:     _post.Write,
			IsArchive: false,
		}

		if err := db.Create(&newPost).Error; err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"filename": fileName,
			"message":  "Yeni gönderi başarıyla paylaşıldı.",
			"post":     newPost,
		})
	}

	return c.JSON(fiber.Map{
		"message": "hata",
	})
}
func (post Post) Delete(c *fiber.Ctx) error {
	isLogin := Helpers.IsLogin(c)
	if isLogin {
		db := Database.DB.Db
		postID := c.Params("id")
		var dPost Models.Post

		if err := db.First(&dPost, "id = ?", postID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Gönderi bulunamadı.",
			})
		}
		id := getID(c)
		if strconv.FormatUint(uint64(dPost.ID), 10) == id {
			if err := db.Delete(&dPost).Error; err != nil {
				return err
			}

			return c.JSON(fiber.Map{
				"message": "Gönderi başarıyla silindi.",
			})
		}
		return c.JSON(fiber.Map{
			"message": "Bu gönderiyi silme yetkiniz yok!!!",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Lütfen önce giriş yapınız.",
	})
}
func (post Post) Archive(c *fiber.Ctx) error {
	isLogin := Helpers.IsLogin(c)
	if isLogin {
		id := getID(c)

		db := Database.DB.Db
		postID := c.Params("id")
		var archivedPost Models.Post
		if postID == id {
			if err := db.First(&archivedPost, postID).Error; err != nil {
				return err
			}
			if err := db.Model(&archivedPost).Where("id=?", postID).Update("is_archive", "1").Error; err != nil {
				return err
			}
			return c.JSON("Başarıyla arşivlendi.")
		}
		return c.JSON("Arşiv için yetkiniz yok!!!")
	}
	return c.JSON("lütfen giriş yapınız!!!")
}
func (post Post) UnArchive(c *fiber.Ctx) error {
	isLogin := Helpers.IsLogin(c)
	if isLogin {
		db := Database.DB.Db
		postID := c.Params("id")
		var archivedPost Models.Post
		id := getID(c)
		if postID == id {
			if err := db.First(&archivedPost, postID).Error; err != nil {
				return err
			}
			if err := db.Model(&archivedPost).Where("id=?", postID).Update("is_archive", "0").Error; err != nil {
				return err
			}
			return c.JSON("Başarıyla arşivden çıkarıldı.")
		}
		return c.JSON("Arşiv için yetkiniz yok!!!")
	}
	return c.JSON("Önce giriş yapınız.")
}

func (post Post) ViewPosts(c *fiber.Ctx) error {
	isLogin := Helpers.IsLogin(c)
	if isLogin {
		db := Database.DB.Db
		var viewPost []Models.Post
		var post []Models.Post
		var users []Models.User
		var follow []Models.Follow
		var follower Models.Follow
		var user Models.User

		id := getID(c)
		db.First(&user, "id=?", id)
		db.Find(&users)
		fmt.Println(getID(c))

		for _, y := range users {
			db.First(&follower, "followed_user_name=? AND follower_user_name=?", y.UserName, user.UserName)
			follow = append(follow, follower)
		}
		for _, a := range follow {
			if err := db.Find(&viewPost, "user_name = ? AND is_archive = ?", a.FollowedUserName, false).Error; err != nil {
				return err
			}
		}
		db.Find(&post, "user_name = ? AND is_archive = ?", user.UserName, false)
		for a := range post {
			viewPost = append(viewPost, post[a])
		}
		return c.JSON(fiber.Map{
			"Postlar": viewPost,
		})
	}

	return c.JSON(fiber.Map{
		"Message": "Önce giriş yapmalısınız!!!",
	})

}

//asdf1213
