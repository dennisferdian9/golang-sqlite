package router

import (
	"net/http"

	"github.com/dennisferdian9/golang-sqlite/controllers"
	"github.com/dennisferdian9/golang-sqlite/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")
	{
		user := api.Group("/user")
		user.Use(middleware.UserMiddleware)
		user.GET("", controllers.GetUsers)
		user.GET(":username", controllers.GetOneUsers)

		user.POST("", controllers.PostUser)
		//  user.GET("/:userID", init.UserCtrl.GetUserById)
		//  user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		//  user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
