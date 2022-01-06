package main

import (
	"gitee.com/snxamdf/http-server/common"
	"gitee.com/snxamdf/http-server/server"
)

func main() {
	common.ColorPrint("Http Server 启动中......\n", common.FontColor.White)
	println()
	common.ColorPrint("免责声明：请不要将该软件做为商业用途，本软件使用过程中造成的损失作者本人概不负责。本软件只做分享学习使用。\n", common.FontColor.Red)
	println()
	common.ColorPrint("软件说明：\n", common.FontColor.Light_blue)
	common.ColorPrint("\t本程序是一个简单易用的网站WEB服务，可用做纯静态文件网站，默认监听80端口。\n", common.FontColor.Light_blue)
	common.ColorPrint("\t本程序不依赖任何组件和程序，纯windows执行程序，如需要linux或mac请联系作者编译。\n", common.FontColor.Light_blue)
	println()
	common.ColorPrint("使用说明：\n", common.FontColor.Light_purple)
	common.ColorPrint("\t1. 双击【http_server.exe】程序或【started.bat】脚本即可启动网站服务\n", common.FontColor.Light_purple)
	common.ColorPrint("\t2. 默认监听80端口，如果想更改端口号，修改started.bat脚本，-port=\"修改端口号\"参数\n", common.FontColor.Light_purple)
	common.ColorPrint("\t\t 端口号修改示例：【http_server -port=\"10088\"】 即监听10088端口\n", common.FontColor.Yellow)
	common.ColorPrint("\t3. 本程序放在指定文件夹做为服务目录\n", common.FontColor.Light_purple)
	println()
	common.ColorPrint("▁▂▃▄▅▆▇▉▉▉▉▉▉▉  程序信息  ▉▉▉▉▉▉▇▆▅▄▃▂▁\n", common.FontColor.Light_yellow)
	common.ColorPrint("作  者：YHY\n", common.FontColor.Light_yellow)
	common.ColorPrint("Email：snxamdf@126.com\n", common.FontColor.Light_yellow)
	common.ColorPrint("Q   Q：122798224\n", common.FontColor.Light_yellow)
	println()
	common.ColorPrint("", common.FontColor.White)
	server.StartHttpServer()
}
