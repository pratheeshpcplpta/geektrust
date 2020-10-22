package color

import "fmt"

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	SuccessColor = "\033[0;32m%s\033[0m"
)

func Info(str string) {
	fmt.Printf(InfoColor, str)
	fmt.Println()
}

func Notice(str string) {
	fmt.Printf(NoticeColor, str)
	fmt.Println()
}

func Warning(str string) {
	fmt.Printf(WarningColor, str)
	fmt.Println()
}

func Error(str string) {
	fmt.Printf(ErrorColor, str)
	fmt.Println()
}

func Debug(str string) {
	fmt.Printf(DebugColor, str)
	fmt.Println()
}

func Success(str string) {
	fmt.Printf(SuccessColor, str)
	fmt.Println()
}
