package service

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/mattn/go-xmlrpc"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"typoraImg/config"
	"typoraImg/consts"
	"typoraImg/utils"
)

func UploadMediaContent(filePath string) (string, error) {
	userConfig := config.UserConfig

	// 读取文件
	var imgBytes []byte
	var readErr error
	if strings.HasPrefix(filePath, consts.HTTPPrefix) {
		imgBytes, readErr = getRemotePic(filePath)
	} else {
		imgBytes, readErr = ioutil.ReadFile(filePath)
	}
	if readErr != nil {
		return "", utils.Error.NewError("读取文件失败: %s", readErr.Error())
	}

	// 文件名 & 文件类型
	fileName, fileType := path.Base(filePath), mimetype.Detect(imgBytes)

	var fileData = xmlrpc.Struct{
		"name": fileName,
		"type": fileType.String(),
		"bits": imgBytes,
	}
	res, err := xmlrpc.Call(
		userConfig.BlogAddress,
		consts.NewMediaObject,
		userConfig.UserAccount,
		userConfig.UserAccount,
		userConfig.AccessToken,
		fileData,
	)
	if err != nil {
		return "", utils.Error.NewError("文件上传失败: %s", err.Error())
	}

	newRes := res.(xmlrpc.Struct)
	if newRes != nil {
		return cast.ToString(newRes["url"]), nil
	}
	return "", utils.Error.NewError("返回数据解析失败！")
}

func getRemotePic(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	return ioutil.ReadAll(res.Body)
}
