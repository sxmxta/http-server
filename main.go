package main

import (
	"flag"
	"http_server/server"
	"time"
)

func sleep(d int64) {
	time.Sleep(time.Duration(d) * time.Second)
}

func main() {
	flag.Parse()
	ColorPrint("Http Server 启动中......\n", FontColor.white)
	sleep(1)
	println()
	ColorPrint("免责声明：请不要将该软件做为商业用途，本软件使用过程中造成的损失作者本人概不负责。本软件只做分享学习使用。\n", FontColor.red)
	println()
	ColorPrint("软件说明：\n", FontColor.light_blue)
	ColorPrint("\t本程序是一个简单易用的网站WEB服务，可用做纯静态文件网站，默认监听80端口。\n", FontColor.light_blue)
	ColorPrint("\t本程序不依赖任何组件和程序，纯windows执行程序，如需要linux或mac请联系作者编译。\n", FontColor.light_blue)
	println()
	ColorPrint("使用说明：\n", FontColor.light_purple)
	ColorPrint("\t1. 双击【http_server.exe】程序或【started.bat】脚本即可启动网站服务\n", FontColor.light_purple)
	ColorPrint("\t2. 默认监听80端口，如果想更改端口号，修改started.bat脚本，-port=\"修改端口号\"参数\n", FontColor.light_purple)
	ColorPrint("\t\t 端口号修改示例：【http_server -port=\"10088\"】 即监听10088端口\n", FontColor.yellow)
	ColorPrint("\t3. 本程序放在指定文件夹做为服务目录\n", FontColor.light_purple)
	println()
	ColorPrint("▁▂▃▄▅▆▇▉▉▉▉▉▉▉  程序信息  ▉▉▉▉▉▉▇▆▅▄▃▂▁\n", FontColor.light_yellow)
	ColorPrint("作者：YHY\n", FontColor.light_yellow)
	ColorPrint("网址：www.snsxm.top\n", FontColor.light_yellow)
	ColorPrint("Q Q：122798224\n", FontColor.light_yellow)
	println()
	ColorPrint("", FontColor.white)
	server.StartHttpServer()
}
