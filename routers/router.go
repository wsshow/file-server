package routers

import (
	"file-server/controllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/", controllers.GetCurDirInfo())
}
