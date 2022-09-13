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

	router.GET("/api/banner", controllers.GetBanners)
	router.GET("/api/banner/:id", controllers.GetBanner)
	router.POST("/api/banner", controllers.CreateBanner)
	router.PATCH("/api/banner/:id", controllers.UpdateBanner)
	router.DELETE("/api/banner/:id", controllers.DeleteBanner)
	
	router.GET("/api/category", controllers.GetCategories)
	router.GET("/api/category/:id", controllers.GetCategory)
	router.POST("/api/category", controllers.CreateCategory)
	router.PATCH("/api/category/:id", controllers.UpdateCategory)
	router.DELETE("/api/category/:id", controllers.DeleteCategory)

	router.GET("/api/customer", controllers.GetCustomers)
	router.GET("/api/customer/:id", controllers.GetCustomer)
	router.POST("/api/customer", controllers.CreateCustomer)
	router.PATCH("/api/customer/:id", controllers.UpdateCustomer)
	router.DELETE("/api/customer/:id", controllers.DeleteCustomer)

	router.GET("/api/mantra", controllers.GetMantras)
	router.GET("/api/mantra/:id", controllers.GetMantra)
	router.POST("/api/mantra", controllers.CreateMantra)
	router.PATCH("/api/mantra/:id", controllers.UpdateMantra)
	router.DELETE("/api/mantra/:id", controllers.DeleteMantra)

	router.GET("/api/media", controllers.GetMedias)
	router.POST("/api/media", controllers.CreateMedia)
	router.DELETE("/api/media/:id", controllers.DeleteMedia)

	router.GET("/api/message", controllers.GetMessages)
	router.GET("/api/message/:id", controllers.GetMessage)
	router.POST("/api/message", controllers.CreateMessage)
	router.PATCH("/api/message/:id", controllers.UpdateMessage)
	router.DELETE("/api/message/:id", controllers.DeleteMessage)

	router.GET("/api/product", controllers.GetProducts)
	router.GET("/api/product/:id", controllers.GetProduct)
	router.POST("/api/product", controllers.CreateProduct)
	router.PATCH("/api/product/:id", controllers.UpdateProduct)
	router.DELETE("/api/product/:id", controllers.DeleteProduct)
	router.GET("/api/printful/products", controllers.PrintfulGetProducts)
	router.GET("/api/printful/products/:id", controllers.PrintfulGetProductById)

	router.GET("/api/page", controllers.GetPages)
	router.GET("/api/page/:id", controllers.GetPage)
	router.GET("/api/page/code/:code", controllers.GetPageByCode)
	router.POST("/api/page", controllers.CreatePage)
	router.PATCH("/api/page/:id", controllers.UpdatePage)
	router.DELETE("/api/page/:id", controllers.DeletePage)

	router.GET("/api/post", controllers.GetPosts)
	router.GET("/api/post/:id", controllers.GetPost)
	router.POST("/api/post", controllers.CreatePost)
	router.PATCH("/api/post/:id", controllers.UpdatePost)
	router.DELETE("/api/post/:id", controllers.DeletePost)
	
	router.GET("/api/subscriber", controllers.GetSubscribers)
	router.GET("/api/subscriber/:id", controllers.GetSubscriber)
	router.POST("/api/subscriber", controllers.CreateSubscriber)
	router.PATCH("/api/subscriber/:id", controllers.UpdateSubscriber)
	router.DELETE("/api/subscriber/:id", controllers.DeleteSubscriber)

	router.POST("/api/upload", controllers.Upload)

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
