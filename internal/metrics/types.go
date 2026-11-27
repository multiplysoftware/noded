package metrics

import "time"

type Metrics struct {
	Time   time.Time     `json:"time"`
	Memory MemoryMetrics `json:"memory"`
}

type MemoryMetrics struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}
