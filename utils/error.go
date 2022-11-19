package utils

import (
	"errors"
	"fmt"
	"os"
)

var Error = &errx{}

type errx struct{}

// ExitWithMsg 打印一条报错消息，并退出 os.Exit(1)
func (e errx) ExitWithMsg(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func (e errx) ExitIf(err error) {
	if err != nil {
		e.ExitWithMsg(err.Error())
	}
}

func (e errx) NewError(format string, a ...any) error {
	return errors.New(fmt.Sprintf(format, a...))
}
