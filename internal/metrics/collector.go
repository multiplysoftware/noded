package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
)

type MetricsCollector struct{}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{}
}

func (c *MetricsCollector) Collect(ctx context.Context) (*Metrics, error) {
	m, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to collect memory metrics: %w", err)
	}

	return &Metrics{
		Timestamp: time.Now().UTC(),
		Memory: MemoryMetrics{
			Total:     m.Total,
			Used:      m.Used,
			Available: m.Available,
		},
	}, nil
}
