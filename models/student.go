package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"primaryKey"`
	Password  string
	FirstName string
	LastName  string
	FacultyID string `gorm:"index"`
}
