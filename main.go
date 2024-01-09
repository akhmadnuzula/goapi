package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"goapi/database"
	"goapi/handlers"
)

func main() {
	// get env
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// koneksi ke db
	database.InitDB()

	// user route
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
	// product route
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProduct)
	router.POST("/products", handlers.CreateProduct)
	router.PUT("/products/:id", handlers.UpdateProduct)
	router.DELETE("/products/:id", handlers.DeleteProduct)
	// auth login
	router.POST("/login", handlers.Login)
	
	// run server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	router.Run(port)
}