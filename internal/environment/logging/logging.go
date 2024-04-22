package logging

import (
	"fmt"
	"log"
)

const errorColor string = "\033[31m"
const infoColor string = "\033[34m"
const successColor string = "\033[32m"
const colorReset string = "\033[0m"
const warningColor string = "\033[33m"

type Logger struct{}

func Info(logEntry string) {

	log.Println(string(infoColor), logEntry, string(colorReset))
}

func Success(logEntry string) {

	log.Println(string(successColor), logEntry, string(colorReset))
}

func Error(format string, a ...interface{}) {

	log.Println(string(errorColor), fmt.Sprintf(format, a...), string(colorReset))
}

func ErrorObject(err error) {
	Error("%s", err)
}

func Fatal(format string, a ...interface{}) {
	log.Fatal(string(errorColor), fmt.Sprintf(format, a...), string(colorReset))
}

func FatalObject(err error) {
	Fatal("%s", err)
}

func Warn(logEntry string) {
	log.Println(string(warningColor), logEntry, string(colorReset))

}
