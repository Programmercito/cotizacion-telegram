package telegram

import (
	"math"

	"cotizacion-telegram/internal/db"
)

type trendInfo struct {
	exists bool
	up     bool
	equal  bool
	diff   float64
}

func buildTrend(current float64, baseline db.Baseline) trendInfo {
	if !baseline.Exists {
		return trendInfo{}
	}
	diff := current - baseline.Value
	return trendInfo{
		exists: true,
		up:     diff > 0,
		equal:  math.Abs(diff) < 0.005,
		diff:   diff,
	}
}
