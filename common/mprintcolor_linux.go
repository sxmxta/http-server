//go:build linux
// +build linux

package common

var (
	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

type Color struct {
	black        int // 黑色
	blue         int // 蓝色
	green        int // 绿色
	cyan         int // 青色
	Red          int // 红色
	Purple       int // 紫色
	Yellow       int // 黄色
	Light_gray   int // 淡灰色（系统默认值）
	Gray         int // 灰色
	Light_blue   int // 亮蓝色
	Light_green  int // 亮绿色
	Light_cyan   int // 亮青色
	Light_red    int // 亮红色
	Light_purple int // 亮紫色
	Light_yellow int // 亮黄色
	White        int // 白色
}

// 输出有颜色的字体
func ColorPrint(s string, i int) {
	println(s)
}

func TestColor() {
	ColorPrint(`红色`, FontColor.Red)
	ColorPrint(`蓝色`, FontColor.blue)
	ColorPrint(`白色`, FontColor.White)
}