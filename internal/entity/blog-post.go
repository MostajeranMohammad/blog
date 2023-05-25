package entity

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	Title    string `gorm:"UNIQUE,NOT NULL"`
	AuthorId uint
	Content  string `gorm:"NOT NULL"`
}
