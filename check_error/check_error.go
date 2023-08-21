package check_error

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
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

func ParserError(err error, ctx *fasthttp.RequestCtx) bool {
	if err != nil {
		fmt.Println("Greška prilikom parsiranja JSON-a:", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Neispravan JSON format")
		return true
	}

	return false
}
