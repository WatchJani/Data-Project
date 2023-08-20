package check_error

import (
	"fmt"
	"log"
)

func ErrorHandler(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func CustomErrorHandler(e error, msg string) {
	if e != nil {
		log.Println(msg, e)
	}
}
