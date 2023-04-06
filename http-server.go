package main

import (
	"embed"
	"gitee.com/snxamdf/http-server/src/entity"
	"gitee.com/snxamdf/http-server/src/gui"
	"gitee.com/snxamdf/http-server/src/server"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types/colors"
)

//go:embed resources
var resources embed.FS

//go:embed libs
var libs embed.FS

func main() {
	inits.Init(&libs, &resources)
	go func() {
		gui.LogsColor(colors.ClRed, "免责声明！\n\t请不要将该软件做为商业用途，本软件使用过程中造成的损失作者本人概不负责。本软件只做分享学习使用。")
		gui.Logs("")
		gui.LogsColor(colors.ClBrown, "软件说明")
		gui.Logs("\t高性能HTTP服务，可用做纯静态网页服务、代理转发、代理拦截，默认监听80端口。")
		gui.Logs("\t跨平台，支持windows、linux、macOS。")
		gui.Logs("")
		gui.LogsColor(colors.ClBlueviolet, "使用说明")
		gui.Logs("\t1. 本程序放在指定文件夹做为服务目录")
		gui.Logs("\t2. 执行【http-server】程序即可启动服务")
		gui.Logs("\t3. 默认监听80端口，端口在【hs.conf.json】配置")
		gui.Logs("\t\t\"server\": { \"ip\": \"0.0.0.0\", \"port\": \"11111\"}")
		gui.Logs("\t4. http代理转发，转发在【hs.conf.json】配置")
		gui.Logs("")
		gui.LogsColor(colors.ClBlue, "▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉")
		gui.Logs("▉\t\t\t作  者：YHY\t\t\t\t\t        ▉")
		gui.Logs("▉\t\t\tEmail：snxamdf@126.com\t\t\t        ▉")
		gui.Logs("▉\t\t\tQ   Q：122798224\t\t\t\t        ▉")
		gui.Logs("▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉▉")
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
