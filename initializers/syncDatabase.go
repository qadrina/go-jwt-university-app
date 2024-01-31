package initializers

import "github.com/qadrina/go-jwt-university-app/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Student{})
}
