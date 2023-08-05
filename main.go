package main

import (
	"log"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/controller"
	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	config := config.GetConfig()
	database.InitDB(config)
	database.DbInstance.CreateTables()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	cityGroup := router.Group("/city")
	cityGroup.POST("", controller.CreateCities)
	cityGroup.GET("", controller.ShowCities)

	theater := router.Group("/theater")
	theater.POST("", controller.AddTheaters)
	theater.GET("", controller.ShowTheaters)
	theater.GET("/city", controller.GetTheatresByCity)
	theater.POST("/addShow", controller.AddShowsInTheatre)
	theater.POST("/show/schedule", controller.AddScreenShowScheduleInTheatre)
	theater.GET("/getShow", controller.GetShowsForTheatre)

	screen := router.Group("/screen")
	screen.POST("", controller.AddScreens)
	screen.GET("", controller.ShowScreens)

	show := router.Group("/show")
	show.POST("", controller.AddShows)
	show.GET("", controller.GetShows)

	auth := router.Group("/auth")
	auth.POST("/register", controller.RegisterUser)
	auth.POST("/login", controller.LoginUser)
	auth.POST("/logout", controller.LogoutUser)

	router.Run(":8000")
	log.Println("Server running on port 8000")
}
