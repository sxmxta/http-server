package main

import (
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/gui"
	"gitee.com/snxamdf/http-server/src/server"
	"time"

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
	lcl.Application.SetTitle("Http Web Server")
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&gui.GUIForm, true)

	t := time.Now()
	msg := t.Format("2006-01-02 15:04:05")
	gui.LogsColor("Http Server 启动中......", colors.ClBlue)
	gui.LogsColor("Http Server 启动时间: "+msg, colors.ClBlue)
	gui.Logs("")
	gui.LogsColor("免责声明：请不要将该软件做为商业用途，本软件使用过程中造成的损失作者本人概不负责。本软件只做分享学习使用。", colors.ClRed)
	gui.Logs("")
	gui.Logs("软件说明：")
	gui.Logs("\t本程序是一个简单易用的网站WEB服务，可用做纯静态网页服务，默认监听80端口。")
	gui.Logs("\t本程序内置依赖，纯 windows/linux 执行程序，如需要linux或mac请联系作者编译。")
	gui.Logs("")
	gui.Logs("使用说明：")
	gui.Logs("\t1. 双击【http-server.exe】程序即可启动服务")
	gui.Logs("\t2. 默认监听80端口，更改端口，hs.conf.json => ")
	gui.Logs("\t\t\"server\": { \"ip\": \"0.0.0.0\", \"port\": \"11111\"}")
	gui.Logs("\t3. 可做http代理转发，代理转发配置，hs.conf.json")
	gui.Logs("\t4. 本程序放在指定文件夹做为服务目录")
	gui.Logs("")
	gui.LogsColor("▁▂▃▄▅▆▇▉▉▉▉▉▉▉  程序信息  ▉▉▉▉▉▉▇▆▅▄▃▂▁", colors.ClBlue)
	gui.LogsColor("作  者：YHY", colors.ClBlue)
	gui.LogsColor("Email：snxamdf@126.com", colors.ClBlue)
	gui.LogsColor("Q   Q：122798224", colors.ClBlue)
	gui.LogsColor("▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉", colors.ClBlue)

	go server.StartHttpServer()
	lcl.Application.Run()
}
