package models

type Faculty struct {
	ID    string `gorm:"primaryKey"`
	Title string
}
