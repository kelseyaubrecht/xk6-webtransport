package wt

import (
	"log"
	"strconv"
	"time"

	"go.k6.io/k6/metrics"
)

func (c *Connection) Write(p []byte) {
	c.requestTimes = append(c.requestTimes, time.Now())
	n, err := c.Stream.Write(p)
	defer c.logWriteMetrics(n)
	if err != nil {
		log.Println("Write error: " + err.Error())
		if n != len(p) {
			log.Println("Wrote n: " + strconv.Itoa(n) + " bytes instead of the expected: " + strconv.Itoa(len(p)))
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
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.WriteCount},
				Value:      1,
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.WriteBytes},
				Value:      float64(n),
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.WriteSize},
				Value:      float64(n),
			},
		},
		Time: now,
	})
}
