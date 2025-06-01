package utils

import (
	"fmt"
)

func ClearConsole() {
	// ANSI escape code to clear screen and move cursor to top-left
	fmt.Print("\033[H\033[2J")
}

type LogColor int

const (
	Black LogColor = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Reset
)

var LogColorStateValue = map[LogColor]string{
	Black:   "\033[30m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
	Reset:   "\033[0m",
}

func (enum LogColor) String() string {
	return LogColorStateValue[enum]
}

func Log(text string, color LogColor) {
	fmt.Printf("%s%s%s ", color, text, Reset)
}
