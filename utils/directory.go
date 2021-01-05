package utils

import (
	"go-svc-bootstrap/global"
	"os"

	"go.uber.org/zap"
)

// @title    PathExists
// @description   check dir exsit or not
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    createDir
// @description   create dirs
// @auth                     （2020/04/05  20:22）
// @param     dirs            string
// @return    err             error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.G_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.G_LOG.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
