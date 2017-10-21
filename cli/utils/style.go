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
