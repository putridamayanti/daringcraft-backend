package main

import (
	"daringcraft-backend/controllers"
	"daringcraft-backend/database"
	"daringcraft-backend/lib"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	if !database.Init() {
		log.Printf("Connected to MongoDB URI: Failure")
		return
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	gin.Logger()

	router.Use(lib.CorsMiddleware())
	router.Use(gin.Logger())

	router.GET("/api/category", controllers.GetCategories)
	router.GET("/api/category/:id", controllers.GetCategory)
	router.POST("/api/category", controllers.CreateCategory)
	router.PATCH("/api/category/:id", controllers.UpdateCategory)
	router.DELETE("/api/category/:id", controllers.DeleteCategory)

	router.GET("/api/product", controllers.GetProducts)
	router.GET("/api/product/:id", controllers.GetProduct)
	router.POST("/api/product", controllers.CreateProduct)
	router.PATCH("/api/product/:id", controllers.UpdateProduct)
	router.DELETE("/api/product/:id", controllers.DeleteProduct)
	router.GET("/api/printful/products", controllers.PrintfulGetProducts)
	router.GET("/api/printful/products/:id", controllers.PrintfulGetProductById)

	router.GET("/api/user", controllers.GetUsers)
	router.GET("/api/user/:id", controllers.GetUser)
	router.POST("/api/user", controllers.CreateUser)
	router.PATCH("/api/user/:id", controllers.UpdateUser)
	router.DELETE("/api/user/:id", controllers.DeleteUser)

	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
