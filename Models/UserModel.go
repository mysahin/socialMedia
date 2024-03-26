package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	LastName      string `json:"last_name"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	SecretAccount bool   `json:"secret_account"`
	Followers     string `json:"followers"`
	FollowerCount int    `json:"follower_count"`
	Bio           string `json:"bio"`
}
