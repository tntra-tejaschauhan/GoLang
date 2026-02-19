package service

import (
	"sync/atomic"
	"time"
)

type Metrics struct {
	TotalRequests int64
	Success       int64
	Failed        int64
	Timeout       int64
	TotalLatency  int64 // nanoseconds
}

var GlobalMetrics = &Metrics{}

func (m *Metrics) Record(duration time.Duration, status string) {

	atomic.AddInt64(&m.TotalRequests, 1)
	atomic.AddInt64(&m.TotalLatency, duration.Nanoseconds())

	switch status {
	case "success":
		atomic.AddInt64(&m.Success, 1)
	case "failed":
		atomic.AddInt64(&m.Failed, 1)
	case "timeout":
		atomic.AddInt64(&m.Timeout, 1)
	}
}

// Latency = how long one request takes to complete.
func (m *Metrics) AverageLatency() time.Duration {
	total := atomic.LoadInt64(&m.TotalRequests)
	if total == 0 {
		return 0
	}
	lat := atomic.LoadInt64(&m.TotalLatency)
	return time.Duration(lat / total)
}
