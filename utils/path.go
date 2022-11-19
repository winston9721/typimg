package utils

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"typoraImg/consts"
)

var Path = &path{}

type path struct{}

func (p path) GetConfigFileFullPath() (string, error) {
	rePath := consts.ConfigFileName
	abPath, err := homedir.Expand(rePath)
	return abPath, err
}

func (p path) GetExecutableAbPath() (string, error) {
	abPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", errors.New("上传命令获取失败，请自行设置")
	}
	return abPath, nil
}
