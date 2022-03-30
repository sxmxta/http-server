package main

import (
	"gitee.com/snxamdf/golcl/tools/winRes"
)

func main() {
	syso()
}

func syso() {
	newSYSO := winRes.NewSYSO()
	newSYSO.RC()
}
