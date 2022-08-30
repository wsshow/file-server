package utils

import (
	"file-server/storage"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func CreatDir(dirPath string) {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Println("CreatDir error:", err)
		return
	}
	err = os.Chmod(dirPath, 0777)
	if err != nil {
		log.Println("Chmod error:", err)
		return
	}
}

func SuitableDisplaySize(size int64) string {
	if size > (1 << 30) {
		return strconv.FormatInt(size>>30, 10) + "GB"
	} else if size > (1 << 20) {
		return strconv.FormatInt(size>>20, 10) + "MB"
	} else if size > (1 << 10) {
		return strconv.FormatInt(size>>10, 10) + "KB"
	} else {
		return strconv.FormatInt(size, 10) + "B"
	}
}

// 获取指定目录下的文件信息(仅当前目录)
func GetFilesInfo(dirPth string) (fis []storage.FileInfo, err error) {
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	fGetFileType := func(cond bool) string {
		if cond {
			return "文件夹"
		} else {
			return "文件"
		}
	}
	fGetSize := func(size int64) string {
		if size == 0 {
			return ""
		} else {
			return SuitableDisplaySize(size)
		}
	}
	var fi storage.FileInfo
	for _, de := range dir {
		fi.FileName = de.Name()
		fi.FileFullPath = filepath.Join(dirPth, de.Name())
		fi.Type = fGetFileType(de.IsDir())
		info, _ := de.Info()
		fi.ModTime = info.ModTime().String()
		fi.Size = fGetSize(info.Size())
		fis = append(fis, fi)
	}
	return fis, nil
}
