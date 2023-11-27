package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}
