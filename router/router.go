package router

import (
	"file-server/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/BackToRootPath", controller.BackToRootPath())
	r.GET("/BackToPrevPath", controller.BackToPrevPath())
	r.GET("/JoinNextPath", controller.JoinNextPath())
	r.GET("/GetMarkdown", controller.GetMarkdown())
	r.GET("/ReloadCurPath", controller.ReloadCurPath())
	r.GET("/GetCurFileName", controller.GetCurFileName())
	r.POST("/UploadFiles", controller.UploadFiles())
	r.POST("/UploadFile", controller.UploadFile())
	r.POST("/DownloadFile", controller.DownloadFile())
	r.POST("/ZipAndDownloadFile", controller.ZipAndDownloadFile())
	r.DELETE("/DeleteFile", controller.DeleteFile())
}
