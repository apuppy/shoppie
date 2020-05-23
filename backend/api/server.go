package api

import (
	"shoppie/api/router"
)

// Listen api server listen and serve
func Listen(bindAddr string) {
	router.HandlerRouter(bindAddr)
}
