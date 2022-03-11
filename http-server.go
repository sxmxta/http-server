package main

import (
	"gitee.com/snxamdf/http-server/src/gui"
	"gitee.com/snxamdf/http-server/src/server"

	"embed"
	"gitee.com/snxamdf/golcl/inits"
	"gitee.com/snxamdf/golcl/lcl"
)

//go:embed resources
var resources embed.FS

//go:embed libs
var libs embed.FS

func main() {
	inits.Init(&libs, &resources)
	lcl.Application.Initialize()
	lcl.Application.SetTitle("MIR客户端")
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&gui.GUIForm, true)

	gui.Logs("Http Server 启动中......")
	gui.Logs("")
	gui.Logs("免责声明：请不要将该软件做为商业用途，本软件使用过程中造成的损失作者本人概不负责。本软件只做分享学习使用。gui.Logs")
	gui.Logs("")
	gui.Logs("软件说明：")
	gui.Logs("\t本程序是一个简单易用的网站WEB服务，可用做纯静态文件网站，默认监听80端口。")
	gui.Logs("\t本程序不依赖任何组件和程序，纯windows执行程序，如需要linux或mac请联系作者编译。")
	gui.Logs("")
	gui.Logs("使用说明：")
	gui.Logs("\t1. 双击【http_server.exe】程序或【started.bat】脚本即可启动网站服务")
	gui.Logs("\t2. 默认监听80端口，如果想更改端口号，修改started.bat脚本，-port=\"修改端口号\"参数")
	gui.Logs("\t\t 端口号修改示例：【http_server -port=\"10088\"】 即监听10088端口")
	gui.Logs("\t3. 本程序放在指定文件夹做为服务目录")
	gui.Logs("")
	gui.Logs("▁▂▃▄▅▆▇▉▉▉▉▉▉▉  程序信息  ▉▉▉▉▉▉▇▆▅▄▃▂▁")
	gui.Logs("作  者：YHY")
	gui.Logs("Email：snxamdf@126.com")
	gui.Logs("Q   Q：122798224")
	gui.Logs("▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉")
	gui.Logs("▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉")

	go server.StartHttpServer()
	lcl.Application.Run()
}
