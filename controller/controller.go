package controller

import (
	"file-server/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var resp utils.Response

func GetCurDirInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		dir, _ := os.Getwd()
		fis, err := utils.GetFilesInfo(dir)
		if err != nil {
			c.JSON(http.StatusOK, resp.Failure().WithDesc("获取当前文件夹信息失败"))
			return
		}
		c.JSON(http.StatusOK, resp.Success(fis))
	}
}
