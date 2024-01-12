package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"goapi/database"
	"goapi/handlers"
	"goapi/middlewares"
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

	// auth login
	router.POST("/login", handlers.Login)
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middlewares.AuthMiddleware())
	{
		// rute user
		userRoute := protectedRoutes.Group("/users")
		{
			userRoute.GET("/", handlers.GetUsers)
			userRoute.GET("/:id", handlers.GetUser)
			userRoute.POST("/", handlers.CreateUser)
			userRoute.PUT("/:id", handlers.UpdateUser)
			userRoute.DELETE("/:id", handlers.DeleteUser)
		}
		// rute product
		productRoute := protectedRoutes.Group("products")
		{
			productRoute.GET("/", handlers.GetProducts)
			productRoute.GET("/:id", handlers.GetProduct)
			productRoute.POST("/", handlers.CreateProduct)
			productRoute.PUT("/:id", handlers.UpdateProduct)
			productRoute.DELETE("/:id", handlers.DeleteProduct)
		}
	}
	
	// run server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	router.Run(port)
}