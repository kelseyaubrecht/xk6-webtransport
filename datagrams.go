package wt

import (
	"log"
	"time"

	"go.k6.io/k6/metrics"
)

func (c *Connection) SendDatagram(p []byte) {
	err := c.Session.SendDatagram(p)
	defer c.logSendDatagramMetrics(len(p))
	if err != nil {
		log.Println("SendDatagram error:", err.Error())
	}
}

func (c *Connection) ReceiveDatagram() []byte {
	p, err := c.Session.ReceiveDatagram(c.Session.Context())
	if err != nil {
		log.Println("ReceiveDatagram error:", err.Error())
	}
	defer c.logRecvDatagramMetrics(len(p))
	return p
}

func (c *Connection) logSendDatagramMetrics(n int) {
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
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.DatagramsSentCount},
				Value:      1,
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.DatagramsSentBytes},
				Value:      float64(n),
			},
		},
		Time: now,
	})
}

func (c *Connection) logRecvDatagramMetrics(n int) {
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
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.DatagramsRecvCount},
				Value:      1,
			},
			{
				Time:       now,
				TimeSeries: metrics.TimeSeries{Metric: c.metrics.DatagramsRecvBytes},
				Value:      float64(n),
			},
		},
		Time: now,
	})
}
