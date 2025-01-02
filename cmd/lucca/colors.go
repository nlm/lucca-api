package main

import (
	"github.com/fatih/color"
)

func Work(s string) string {
	return color.RedString(s)
}

func Leave(s string) string {
	return color.HiBlueString("-- %s --", s)
}

func Title(s string) string {
	return color.New(color.BgHiGreen, color.FgBlack, color.Bold).Sprint(s)
}

func Titlef(format string, a ...any) string {
	return color.New(color.BgHiGreen, color.FgBlack, color.Bold).Sprintf(format, a...)
}

func Error(s string) string {
	return color.New(color.BgHiRed, color.FgBlack, color.Bold).Sprint(s)
}

func Info(s string) string {
	return color.New(color.BgHiBlue, color.FgBlack, color.Bold).Sprint(s)
}
