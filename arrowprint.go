package arrowprint

import (
	"fmt"

	"github.com/mattn/go-colorable"
)

// helper functions
func printColor0(cc int, msg string, args ...any) {
	fmt.Fprintf(
		colorable.NewColorableStdout(),
		"\033[3%dm==>\033[0m %s\n",
		cc,
		fmt.Sprintf(msg, args...))
}

func printColor1(cc int, msg string, args ...any) {
	fmt.Fprintf(
		colorable.NewColorableStdout(),
		"\033[3%dm  ->\033[0m %s\n",
		cc,
		fmt.Sprintf(msg, args...))
}

func printColon(cc int, msg string, args ...any) {
	fmt.Fprintf(
		colorable.NewColorableStdout(),
		"\033[3%d;1m ::\033[0m\033[1m %s\033[0m\n",
		cc,
		fmt.Sprintf(msg, args...))
}

// colon print
// :: Message
func InfoC(msg string, args ...any) {
	printColon(4, msg, args...)
}
func SucC(msg string, args ...any) {
	printColon(2, msg, args...)
}
func WarnC(msg string, args ...any) {
	printColon(3, msg, args...)
}
func ErrC(msg string, args ...any) {
	printColon(1, msg, args...)
}

// level 0 print
// => message
func Info0(msg string, args ...any) {
	printColor0(4, msg, args...)
}
func Suc0(msg string, args ...any) {
	printColor0(2, msg, args...)
}
func Warn0(msg string, args ...any) {
	printColor0(3, msg, args...)
}
func Err0(msg string, args ...any) {
	printColor0(1, msg, args...)
}

// level 1 print
//   -> message
func Info1(msg string, args ...any) {
	printColor1(4, msg, args...)
}
func Suc1(msg string, args ...any) {
	printColor1(2, msg, args...)
}
func Warn1(msg string, args ...any) {
	printColor1(3, msg, args...)
}
func Err1(msg string, args ...any) {
	printColor1(1, msg, args...)
}
