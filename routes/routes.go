// File: routes/user_routes.go
package routes

import (
	"apis/controller"
	"apis/dao"
	"apis/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	srv := service.NewService(dao.NewMongoDao())
	ctr := controller.NewController(srv)
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", ctr.GetUsers)

		userRoutes.GET("/:name", ctr.GetUserDetails)

		userRoutes.POST("/", ctr.CreateUser)

		userRoutes.DELETE("/:name", ctr.DeleteUser)

		userRoutes.PUT("/inc/hell/:name", ctr.IncHell)
		userRoutes.PUT("/inc/heaven/:name", ctr.IncHeaven)

		userRoutes.PUT("/dec/hell/:name", ctr.DecHell)
		userRoutes.PUT("/dec/heaven/:name", ctr.DecHeaven)

	}
}
