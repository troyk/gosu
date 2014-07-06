package gosu

import (
	"fmt"
	"github.com/mgutz/ansi"
	"runtime"
)

var cyan = ansi.ColorFunc("cyan")
var red = ansi.ColorFunc("red+b")
var redInverse = ansi.ColorFunc("white:red")
var gray = ansi.ColorFunc("black+h")
var magenta = ansi.ColorFunc("magenta+h")

func init() {
	if runtime.GOOS == "windows" {
		ansi.DisableColors(true)
	}
}

// Debugf writes a debug statement to stdout.
func Debugf(group string, format string, any ...interface{}) {
	fmt.Print(gray(group) + " ")
	fmt.Printf(gray(format), any...)
}

// Infof writes an info statement to stdout.
func Infof(group string, format string, any ...interface{}) {
	fmt.Print(cyan(group) + " ")
	fmt.Printf(format, any...)
}

// Errorf writes an error statement to stdout.
func Errorf(group string, format string, any ...interface{}) {
	fmt.Errorf(red(group) + " ")
	fmt.Errorf(red(format), any...)
}

// Panicf writes an error statement to stdout.
func Panicf(group string, format string, any ...interface{}) {
	fmt.Printf(redInverse(group) + " ")
	fmt.Printf(redInverse(format), any...)
}