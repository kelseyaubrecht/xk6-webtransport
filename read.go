package wt

import (
	"io"
	"log"
	"time"

	"go.k6.io/k6/metrics"
)

func (c *Connection) ReadAll() []byte {
	rsp, err := io.ReadAll(c.activeStream)
	defer c.logReadMetrics(len(rsp))

	if err != nil {
		log.Println("Read error:", err.Error())
	}

	return rsp
}

func (c *Connection) ReadFull(expectedReadLength int) []byte {
	rsp := make([]byte, expectedReadLength)
	n, err := io.ReadFull(c.activeStream, rsp)
	defer c.logReadMetrics(n)

	if err != nil {
		log.Println("Read error:", err.Error())
		if n != expectedReadLength {
			log.Printf("Read n: %d does not match the expected length of: %d\n", n, expectedReadLength)
		}
	}

	return rsp
}

func (c *Connection) ReadAtLeast(maxReadLength int, minReadLength int) []byte {
	rsp := make([]byte, maxReadLength)
	n, err := io.ReadAtLeast(c.activeStream, rsp, minReadLength)
	defer c.logReadMetrics(n)

	if err != nil {
		log.Println("Read error:", err.Error())
		if n < minReadLength {
			log.Printf("Read n: %d is smaller than expected minimum: %d\n", n, minReadLength)
		}
	}

	return rsp
}

func (c *Connection) logReadMetrics(n int) {
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
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.StreamsReadCount},
				Value:      1,
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.StreamsReadBytes},
				Value:      float64(n),
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.StreamsReadSize},
				Value:      float64(n),
			},
		},
		Time: now,
	})
}
