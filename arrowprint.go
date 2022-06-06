package arrowprint

import "fmt"

func printColor0(cc int, msg string, args ...interface{}) {
	fmt.Printf("\033[3%dm==>\033[0m %s\n", cc, fmt.Sprintf(msg, args...))
}

func printColor1(cc int, msg string, args ...interface{}) {
	fmt.Printf("\033[3%dm  ->\033[0m %s\n", cc, fmt.Sprintf(msg, args...))
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

func Suc1(msg string, args ...interface{}) {
	printColor1(2, msg, args...)
}

func Warn1(msg string, args ...interface{}) {
	printColor1(3, msg, args...)
}

func Err1(msg string, args ...interface{}) {
	printColor1(1, msg, args...)
}
