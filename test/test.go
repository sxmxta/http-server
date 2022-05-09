package test

import (
	"gitee.com/snxamdf/golcl/tools/winRes"
	"testing"
)

func testSyso(t *testing.T) {
	syso()
}

func syso() {
	newSYSO := winRes.NewSYSO()
	newSYSO.RC()
}

func testLogs(t *testing.T) {
}
