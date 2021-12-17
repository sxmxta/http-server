package modules

import (
	"fmt"
	"http_server/server"
)

func list(ctx *server.Context) {
	fmt.Println(ctx.Get("name"))
}

func test(ctx *server.Context) {
	fmt.Println(ctx.Get("name"))
	ctx.Write([]byte("asdf"))
}
