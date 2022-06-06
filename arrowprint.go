package arrowprint

import "fmt"

func printColor0(cc int, msg string, args ...interface{}) {
	fmt.Printf("\033[3%dm==>\033[0m %s\n", cc, fmt.Sprintf(msg, args...))
}

func printColor1(cc int, msg string, args ...interface{}) {
	fmt.Printf("\033[3%dm  ->\033[0m %s\n", cc, fmt.Sprintf(msg, args...))
}

func printColon(cc int, msg string, args ...interface{}) {
	fmt.Printf("\033[3%dm ::\033[0m %s\n", cc, fmt.Sprintf(msg, args...))
}

func Suc0(msg string, args ...interface{}) {
	printColor0(2, msg, args...)
}
func Warn0(msg string, args ...interface{}) {
	printColor0(3, msg, args...)
}
func Err0(msg string, args ...interface{}) {
	printColor0(1, msg, args...)
}
func Info0(msg string, args ...interface{}) {
	printColor0(4, msg, args...)
}

func Suc1(msg string, args ...interface{}) {
	printColor1(2, msg, args...)
}
func Warn1(msg string, args ...interface{}) {
	printColor1(3, msg, args...)
}
func Err1(msg string, args ...interface{}) {
	printColor1(1, msg, args...)
}
func Info1(msg string, args ...interface{}) {
	printColor1(4, msg, args...)
}

func SucC(msg string, args ...interface{}) {
	printColon(2, msg, args...)
}
func WarnC(msg string, args ...interface{}) {
	printColon(3, msg, args...)
}
func ErrC(msg string, args ...interface{}) {
	printColon(1, msg, args...)
}
func InfoC(msg string, args ...interface{}) {
	printColon(4, msg, args...)
}
