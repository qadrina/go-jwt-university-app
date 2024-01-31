package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qadrina/go-jwt-university-app/controllers"
	"github.com/qadrina/go-jwt-university-app/initializers"
	"github.com/qadrina/go-jwt-university-app/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	//http.HandleFunc("/", controllers.Welcome)
	r.GET("/", controllers.Welcome)
	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/api/students", controllers.GetAllStudents)
	r.Run()
}
