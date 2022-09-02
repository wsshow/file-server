package utils

import (
	"file-server/storage"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func IsDir(filePath string) (bool, error) {
	d, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return d.IsDir(), nil
}

func CreatDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(dirPath, 0777)
	if err != nil {
		return err
	}
	return nil
}

func ReadAll(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func RemoveDir(filePath string) error {
	return os.RemoveAll(filePath)
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
		fi.ModTime = info.ModTime().Format("2006-01-02 15:04:05")
		fi.Size = fGetSize(info.Size())
		fis = append(fis, fi)
	}
	return fis, nil
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:8")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
