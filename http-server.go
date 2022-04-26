package main

import (
	"embed"
	"gitee.com/snxamdf/golcl/inits"
	"gitee.com/snxamdf/golcl/lcl"
	"gitee.com/snxamdf/golcl/lcl/types/colors"
	"gitee.com/snxamdf/http-server/src/entity"
	"gitee.com/snxamdf/http-server/src/gui"
	"gitee.com/snxamdf/http-server/src/server"
)

//go:embed resources
var resources embed.FS

//go:embed libs
var libs embed.FS

func main() {
	inits.Init(&libs, &resources)
	go func() {
		gui.LogsColor(colors.ClRed, "免责声明：请不要将该软件做为商业用途，本软件使用过程中造成的损失作者本人概不负责。本软件只做分享学习使用。")
		gui.Logs("")
		gui.Logs("软件说明")
		gui.Logs("\t本程序是一个简单易用的网站WEB服务，可用做纯静态网页服务，默认监听80端口。")
		gui.Logs("\t本程序内置依赖，纯 windows/linux 执行程序，如需要linux或mac请联系作者编译。")
		gui.Logs("")
		gui.Logs("使用说明")
		gui.Logs("\t1. 双击【http-server.exe】程序即可启动服务")
		gui.Logs("\t2. 默认监听80端口，更改端口，hs.conf.json")
		gui.Logs("\t\t\"server\": { \"ip\": \"0.0.0.0\", \"port\": \"11111\"}")
		gui.Logs("\t3. 可做http代理转发，代理转发配置，hs.conf.json")
		gui.Logs("\t4. 本程序放在指定文件夹做为服务目录")
		gui.Logs("")
		gui.LogsColor(colors.ClBlue, "▁▂▃▄▅▆▇▉▉▉▉▉▉▉  程序信息  ▉▉▉▉▉▉▇▆▅▄▃▂▁")
		gui.Logs("作  者：YHY")
		gui.Logs("Email：snxamdf@126.com")
		gui.Logs("Q   Q：122798224")
		gui.Logs("开发语言：Golang")
		gui.Logs("▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉")
		gui.LogsColor(colors.ClDarkblue, "-------------------------------- 以下服务日志 --------------------------------")
		go func() {
			for {
				//如果多路的话需要 for select 配合 使用
				//单通道只需要for即可
				select {
				case msg, ok := <-entity.GlobalLogMessageChan:
					if ok {
						if msg.Type == 0 {
							gui.Logs(msg.Message...)
						} else if msg.Type == 1 {
							m := ""
							for _, v := range msg.Message {
								m += v
							}
							gui.LogsColor(msg.Color, m)
						} else if msg.Type == 2 {
							gui.LogsTime(msg.Message...)
						} else if msg.Type == 3 {
							gui.LogsProxyTime(msg.Message...)
						} else if msg.Type == 4 {
							gui.LogsStaticTime(msg.Message...)
						}
					}
				}
			}
		}()
		if entity.AppInitSuccess {
			server.StartHttpServer()
		}
	}()
	lcl.RunApp(&gui.GUIForm)
}
