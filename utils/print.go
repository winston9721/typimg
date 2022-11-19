package utils

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
	"typoraImg/consts"
)

type IPrint interface {
	ColorPrint(message, color string)
	ColorPrintln(message, color string)
}

type unixPrint struct{}

func (p unixPrint) ColorPrint(message, color string) {
	_, _ = fmt.Fprint(os.Stdout, ansi.Color(message, color))
}

func (p unixPrint) ColorPrintln(message, color string) {
	_, _ = fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}

type dosPrint struct{}

func (p dosPrint) ColorPrint(message, color string) {
	fmt.Printf(message)
}

func (p dosPrint) ColorPrintln(message, color string) {
	fmt.Println(message)
}

func NewPrint(s string) IPrint {
	switch s {
	case consts.OSWindows:
		return &dosPrint{}
	default:
		return &unixPrint{}
	}
}
