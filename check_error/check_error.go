package check_error

import (
	"fmt"
)

func ErrorHandler(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
