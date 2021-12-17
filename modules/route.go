package modules

import "http_server/server"

func init() {
	server.RegisterRoute("/list", list)
	server.RegisterRoute("/add", add)
}
