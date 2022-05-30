package arrowprint

import "fmt"

func printColor0(cc int, msg string) {
	fmt.Printf("\033[3%dm==> %s\n", cc, msg)
}

func printColor1(cc int, msg string) {
	fmt.Printf("\033[3%dm  -> %s\n", cc, msg)
}

func Suc0(msg string) {
	printColor0(2, msg)
}

func Warn0(msg string) {
	printColor0(3, msg)
}

func Err0(msg string) {
	printColor0(1, msg)
}

func Suc1(msg string) {
	printColor1(2, msg)
}

func Warn1(msg string) {
	printColor1(3, msg)
}

func Err1(msg string) {
	printColor1(1, msg)
}
