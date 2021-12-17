package modules

import "http_server/server"

func init() {
	server.RegisterRoute("/paymonthly/list", list)
	server.RegisterRoute("/paymonthly/add", add)
	server.RegisterRoute("/paymonthly/delete", delete)
}
