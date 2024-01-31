package main

import "github.com/qadrina/go-jwt-university-app/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.SyncDatabase()
}
