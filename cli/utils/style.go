package utils

import "github.com/ttacon/chalk"

type _Style struct{}

var Style _Style

func (_ *_Style) Title(title string) string {
	return chalk.Bold.TextStyle(chalk.Cyan.Color(title))
}

func (_ *_Style) Bold(text string) string {
	return chalk.Bold.TextStyle(text)
}

func (_ *_Style) Success(text string) string {
	return chalk.Bold.TextStyle(chalk.Green.Color(text))
}

func (_ *_Style) Fail(text string) string {
	return chalk.Bold.TextStyle(chalk.Red.Color(text))
}

func (_ *_Style) Disabled(text string) string {
	return chalk.Dim.TextStyle(text)
}
