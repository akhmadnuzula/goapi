package main

import (
	"fmt"
	"goapi/database"
	"goapi/docs"
	"goapi/handlers"
	"goapi/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// "github.com/swaggo/gin-swagger/swaggerFiles"
// docs "goapi/docs"

// @title Your API Title
// @version 1.0
// @description Your API Description
// @termsOfService http://swagger.io/terms/
// @contact.email Your Contact Email
// @license.name Your License Name
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /v1
func main() {
	// get env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	// koneksi ke db
	database.InitDB()

	// Tambahkan rute untuk Swagger Docs
	docs.SwaggerInfo.Title = "Swagger GOAPI"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
