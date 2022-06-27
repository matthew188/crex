package exsim

import (
	"testing"

	. "github.com/matthew188/crex"
)

func TestCalcPnl(t *testing.T) {
	size := 50.0
	entryPrice := 10351.5
	exitPrice := 10348.5
	pnl, pnlUsd := CalcPnl(Buy, size, entryPrice, exitPrice)
	t.Logf("pnl: %.8f", pnl)
	t.Logf("pnlUsd: %.8f", pnlUsd)
}
