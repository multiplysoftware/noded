package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/multiplysoftware/handyman/noded/internal/metrics"
)

func main() {
	l := slog.Default()

	c := metrics.NewMetricsCollector()

	ms, err := c.Collect(context.Background())
	if err != nil {
		l.Error("failed to collect metrics", slog.String("error", err.Error()))
		os.Exit(1)
	}

	l.Info("get metrics", slog.Any("metrics", ms))
}
