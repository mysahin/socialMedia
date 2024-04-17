package Router

import (
	"github.com/gofiber/fiber/v2"
	"socialMedia/Controllers"
)

func Routes() *fiber.App {
	r := fiber.New()
	r.Get("/ana-sayfa", Controllers.Post{}.ViewPosts)
	r.Post("/login", Controllers.Login{}.SignIn)
	r.Post("/sign-up", Controllers.Login{}.SignUp)
	r.Get("/logout", Controllers.Login{}.SignOut)
	r.Post("/share", Controllers.Post{}.Share)
	r.Post("/delete-post/:id", Controllers.Post{}.Delete)
	r.Post("/archive/:id", Controllers.Post{}.Archive)
	r.Post("/un-archive/:id", Controllers.Post{}.UnArchive)
	r.Post("/follow/:id", Controllers.User{}.Follow)
	r.Post("/secret", Controllers.User{}.DoSecret)
	r.Post("/un-secret", Controllers.User{}.DoPublic)
	r.Get("/list", Controllers.ListFiles)
	r.Get("/show/:filename", Controllers.ShowFile)
	r.Post("/delete/:filename", Controllers.DeleteFile)

	return r
}
