package Models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	FileName string `json:"file_name"`
}
