package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserName string
	Body string
}