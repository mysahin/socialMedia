package Router

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"socialMedia/Controllers"
)

func Routes(uploader *s3manager.Uploader, downloader *s3.S3, bucketName string) *fiber.App {
	r := fiber.New()

	// Ana sayfa
	r.Get("/ana-sayfa", Controllers.Post{}.ViewPosts)

	// Giriş ve kayıt işlemleri
	r.Post("/login", Controllers.Login{}.SignIn)
	r.Post("/sign-up", Controllers.Login{}.SignUp)
	r.Get("/logout", Controllers.Login{}.SignOut)

	// Paylaşım işlemleri
	r.Post("/share", Controllers.Post{}.Share)
	r.Post("/delete-post/:id", Controllers.Post{}.Delete)
	r.Post("/archive/:id", Controllers.Post{}.Archive)
	r.Post("/un-archive/:id", Controllers.Post{}.UnArchive)

	// Takip işlemleri
	r.Post("/follow/:id", Controllers.User{}.Follow)

	// Dosya işlemleri
	fileController := Controllers.NewFileController(uploader, downloader, bucketName)
	r.Post("/upload", fileController.UploadFile)
	r.Get("/list", fileController.ListFiles)
	r.Get("/show/:filename", fileController.ShowFile)
	r.Post("/delete/:filename", fileController.DeleteFile)

	return r
}
