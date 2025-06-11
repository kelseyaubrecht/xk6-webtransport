package wt

import (
	"log"
	"strconv"
	"time"

	"go.k6.io/k6/metrics"
)

func (c *Connection) Write(p []byte) {
	n, err := c.activeStream.Write(p)
	defer c.logWriteMetrics(n)
	if err != nil {
		log.Println("Write error:", err.Error())
		if n != len(p) {
			log.Printf("Wrote n: %q bytes instead of the expected: %q \n", strconv.Itoa(n), strconv.Itoa(len(p)))
		}
	}
}

func (c *Connection) logWriteMetrics(n int) {
	state := c.vu.State()
	ctx := c.vu.Context()
	if state == nil || ctx == nil {
		return
	}

	now := time.Now()

	metrics.PushIfNotDone(ctx, state.Samples, metrics.ConnectedSamples{
		Samples: []metrics.Sample{
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.StreamsWriteCount},
				Value:      1,
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.StreamsWriteBytes},
				Value:      float64(n),
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.StreamsWriteSize},
				Value:      float64(n),
			},
		},
		Time: now,
	})
}
