package service

import (
	"errors"
	"github.com/mattn/go-xmlrpc"
	"github.com/spf13/cast"
	"typoraImg/consts"
	"typoraImg/form"
)

func LoginBlog(blogInfo *form.BlogInfo) error {
	res, err := xmlrpc.Call(
		blogInfo.BlogAddress,
		consts.GetUsersBlogs,
		blogInfo.UserAccount,
		blogInfo.UserAccount,
		blogInfo.AccessToken,
	)
	if err != nil {
		return errors.New("博客园登录失败，请仔细确认用户名和密码是否正确！")
	}

	for _, item := range res.(xmlrpc.Array) {
		for k, v := range item.(xmlrpc.Struct) {
			switch k {
			case "blogName":
				blogInfo.UserName = cast.ToString(v)
			}
		}
	}
	return nil
}
