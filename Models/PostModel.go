package Models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserName  string `json:"user_name"`
	Write     string `json:"write"`
	IsArchive bool   `json:"is_archive"`
	FileId    uint   `json:"file_id"`
}
