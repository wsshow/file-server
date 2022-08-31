package controller

import (
	"file-server/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var resp utils.Response
var historyOpPath = utils.NewStack()

func GetCurDirInfo(curPath string) utils.Response {
	fis, err := utils.GetFilesInfo(curPath)
	if err != nil {
		return resp.Failure().WithDesc("获取当前文件夹信息失败")
	}
	return resp.Success(fis)
}

func JoinNextPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		curpath, ok := c.GetQuery("curpath")
		if !ok {
			c.JSON(http.StatusOK, resp.Failure().WithDesc("路径获取失败"))
			return
		}
		if ok, err := utils.IsDir(curpath); err != nil || !ok {
			c.JSON(http.StatusOK, resp.Failure().WithDesc("无效的路径"))
			return
		}
		historyOpPath.Push(curpath)
		c.JSON(http.StatusOK, GetCurDirInfo(curpath))
	}
}

func BackToPrevPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		lastPath := historyOpPath.Pop()
		if lastPath == nil {
			dir, _ := os.Getwd()
			lastPath = dir
		}
		if ok, err := utils.IsDir(lastPath.(string)); err != nil || !ok {
			c.JSON(http.StatusOK, resp.Failure().WithDesc("无效的路径"))
			return
		}
		c.JSON(http.StatusOK, GetCurDirInfo(lastPath.(string)))
	}
}

func BackToRootPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		dir, _ := os.Getwd()
		c.JSON(http.StatusOK, GetCurDirInfo(dir))
	}
}
