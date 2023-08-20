package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/valyala/fasthttp"
)

type Middleware func(fasthttp.RequestHandler) fasthttp.RequestHandler

func chainingMiddleware(h fasthttp.RequestHandler, m ...Middleware) fasthttp.RequestHandler {
	if len(m) < 1 {
		return h
	}

	wrappedHandler := h
	for i := len(m) - 1; i >= 0; i-- {
		wrappedHandler = m[i](wrappedHandler)
	}

	return wrappedHandler
}

func CORSMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		origin := string(ctx.Request.Header.Peek("Origin"))
		ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
		if string(ctx.Method()) == "OPTIONS" {
			ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
			ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,POST")
			ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		} else {
			next(ctx)
		}
	}
}

func panicRecovery(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Error(fmt.Sprintf(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next(ctx)
	}
}

func headerMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Content-Type", "application/json")
		next(ctx)
	}
}

func loggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("Before handler is executed")
		ctx.WriteString("Adding response via middleware\n")
		log.Println(string(ctx.Request.URI().Path()))
		next(ctx)
		fmt.Println("After handler is executed")
	}
}
