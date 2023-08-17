package routes

import (
	"github.com/erabxes/golang-jwt-project/middleware"

	controller "github.com/erabxes/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/user", controller.GetUsers())
	incomingRoutes.GET("/uers/:user_id", controller.GetUser())

}
