package test

import (
	"github.com/energye/golcl/tools/winRes"
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
