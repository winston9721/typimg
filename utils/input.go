package utils

import (
	"runtime"
	"strings"
	"typoraImg/consts"
)

var Input = &input{}

type input struct{}

func (i input) Format(s string) string {
	if runtime.GOOS == consts.OSWindows {
		return strings.Replace(s, "\r\n", "", -1)
	} else {
		return strings.Replace(s, "\n", "", -1)
	}
}
