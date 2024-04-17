package Controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var region string = "eu-north-1"
var accessKey string = "AKIAQ3EGTZABT7PRWHUS"
var secretKey string = "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"

var uploader *s3manager.Uploader
var downloader *s3.S3

var bucketName string = "social-media-mysahin"

func UploadFile(c *fiber.Ctx) (*string, error) {
	var name string
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File["files"]

	for _, file := range files {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		filename := fixFileName(fileHeader.Filename)
		name = filename

		_, err = saveFile(f, filename)
		if err != nil {
			return nil, err
		}
	}

	return &name, nil
}

func fixFileName(filename string) string {
	// Türkçe karakterleri ingilizce karakterlere dönüştür
	replacer := strings.NewReplacer("ı", "i", "ğ", "g", "ü", "u", "ş", "s", "ö", "o", "ç", "c", "İ", "I", "Ğ", "G", "Ü", "U", "Ş", "S", "Ö", "O", "Ç", "C")
	filename = replacer.Replace(filename)

	// Boşlukları kaldır
	filename = strings.ReplaceAll(filename, " ", "")

	return filename
}

func saveFile(fileReader io.Reader, filename string) (string, error) {
	// Upload the file to S3 using the fileReader
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)

	return url, nil
}

func ListFiles(c *fiber.Ctx) error {
	resp, err := downloader.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var filenames []string
	for _, item := range resp.Contents {
		filenames = append(filenames, *item.Key)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"filenames": filenames})
}

func ShowFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	obj, err := downloader.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer obj.Body.Close()

	// Determine the content type based on the file extension
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType == "" {
		// If the content type is not recognized, default to octet-stream
		contentType = "application/octet-stream"
	}

	// Set the content type header
	c.Set("Content-Type", contentType)

	content, err := io.ReadAll(obj.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).Send(content)
}

func DeleteFile(c *fiber.Ctx) error {
	filename := c.Params("filename")

	// Dosyayı S3'den sil
	_, err := downloader.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Printf("Dosya '%s' başarıyla silindi.\n", filename)

	return c.SendStatus(http.StatusOK)
}
