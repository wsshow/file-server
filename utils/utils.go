package utils

import (
	"log"
	"os"
	"path/filepath"
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

// 获取指定目录下的文件和目录(仅当前目录)
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	for _, fi := range dir {
		fullPath := filepath.Join(dirPth, fi.Name())
		if fi.IsDir() {
			dirs = append(dirs, fullPath)
		} else {
			files = append(files, fullPath)
		}
	}
	return files, dirs, nil
}
