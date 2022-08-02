package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stheven26/technical_test/controllers"
)

var (
	router = gin.Default()
)

func Route() {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/message", controllers.OneOnOneMessage)
	router.POST("/group", controllers.CreateGroup)
}

func StartApplication() {
	Route()
	router.Run(":8000")
}
