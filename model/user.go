package model

import "github.com/jinzhu/gorm"

// User represents the user schema for the database
type User struct {
	gorm.Model
	Email    string
	Password string
}

// type WebhookUser struct {
// 	Email string
// }
