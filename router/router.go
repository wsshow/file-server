package router

import (
	"file-server/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/BackToRootPath", controller.BackToRootPath())
	r.GET("/BackToPrevPath", controller.BackToPrevPath())
	r.GET("/JoinNextPath", controller.JoinNextPath())
	r.GET("/ReloadCurPath", controller.ReloadCurPath())
	r.POST("/UploadFiles", controller.UploadFiles())
	r.POST("/UploadFile", controller.UploadFile())
}
