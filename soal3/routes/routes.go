package routes

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stheven26/technical_test/controllers"
	"github.com/stheven26/technical_test/globals"
)

var (
	router = gin.Default()
)

func Route() {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	private := router.Group("/api/v1")

	//middleware authorization
	private.Use(jwt.Auth(globals.Key))
	private.POST("/message", controllers.OneOnOneMessage)
	private.POST("/group", controllers.CreateGroup)
}

func StartApplication() {
	Route()
	router.Run(":8000")
}
