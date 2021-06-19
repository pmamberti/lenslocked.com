package models

import "github.com/jinzhu/gorm"

// Gallery represents a gallery resource
type Gallery struct {
	gorm.Model
	UserID uint   `gorm:"not_null;index"`
	Title  string `gorm:"not_null"`
}
