package Routes

import (
	"Project/Handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", Handler.CreateUser)
}
