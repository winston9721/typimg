package config

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"typoraImg/form"
	"typoraImg/utils"
)

var UserConfig form.BlogInfo

func NewConfig() (*viper.Viper, error) {
	config := viper.New()
	filePath, err := utils.Path.GetConfigFileFullPath()
	if err == nil {
		config.SetConfigFile(filePath)
	}
	return config, err
}

func UpCreateConfig(blogInfo form.BlogInfo) error {
	config, err := NewConfig()
	if err != nil {
		return utils.Error.NewError("配置项初始化失败01: %s", err.Error())
	}
	info, err := json.Marshal(blogInfo)
	if err != nil {
		return utils.Error.NewError("配置项初始化失败02: %s", err.Error())
	}
	if err = config.ReadConfig(bytes.NewReader(info)); err != nil {
		return utils.Error.NewError("配置文件保存失败01: %s", err.Error())
	}

	if err = config.WriteConfig(); err != nil {
		return utils.Error.NewError("配置文件保存失败02: %s", err.Error())
	}
	return nil
}

func LoadConfig() error {
	config, err := NewConfig()
	if err != nil {
		return utils.Error.NewError("配置项初始化失败: %s", err.Error())
	}
	if err = config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return utils.Error.NewError("配置文件不存在，请执行初始化[init]命令")
		} else {
			return utils.Error.NewError("配置文件读取失败: %s", err.Error())
		}
	}
	if err = config.Unmarshal(&UserConfig); err != nil {
		return utils.Error.NewError("配置文件加载失败: %s", err.Error())
	}
	return nil
}
