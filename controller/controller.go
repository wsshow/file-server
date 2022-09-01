package controller

import (
	"file-server/utils"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var resp utils.Response
var historyOpPath = utils.NewStack()
var currentPath string = "."

func GetCurDirInfo(curPath string) utils.Response {
	fis, err := utils.GetFilesInfo(curPath)
	if err != nil {
		return resp.Failure().WithDesc("获取当前文件夹信息失败")
	}
	currentPath = curPath
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

func ReloadCurPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, GetCurDirInfo(currentPath))
	}
}

func BackToPrevPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		lastPath := historyOpPath.Pop()
		for lastPath != nil && lastPath.(string) == currentPath {
			lastPath = historyOpPath.Pop()
		}
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

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, resp.Failure().WithDesc(err.Error()))
			return
		}
		dir, _ := os.Getwd()
		dst := path.Join(dir, file.Filename)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusOK, resp.Failure().WithDesc(err.Error()))
			return
		}
		c.JSON(http.StatusOK, resp.Success(nil))
	}
}

func UploadFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		var urls []string
		form, _ := c.MultipartForm()
		files := form.File["files"]
		dir, _ := os.Getwd()
		for _, file := range files {
			fmt.Println(file.Filename)
			dst := path.Join(dir, file.Filename)
			urls = append(urls, dst)
			c.SaveUploadedFile(file, dst)
		}
		fmt.Println(urls)
		c.JSON(http.StatusOK, GetCurDirInfo(dir))
	}
}
