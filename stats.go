package wt

import (
	"errors"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type WTMetrics struct {
	StreamsWriteCount *metrics.Metric
	StreamsWriteBytes *metrics.Metric
	StreamsWriteSize  *metrics.Metric

	StreamsReadCount *metrics.Metric
	StreamsReadBytes *metrics.Metric
	StreamsReadSize  *metrics.Metric

	StreamsTotal *metrics.Metric

	DatagramsSentCount *metrics.Metric
	DatagramsSentBytes *metrics.Metric

	DatagramsRecvCount *metrics.Metric
	DatagramsRecvBytes *metrics.Metric
}

func registerMetrics(vu modules.VU) (WTMetrics, error) {
	var err error
	registry := vu.InitEnv().Registry
	wtMetrics := WTMetrics{}

	if wtMetrics.StreamsWriteCount, err = registry.NewMetric(
		"webtransport_write_count", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.StreamsWriteBytes, err = registry.NewMetric(
		"webtransport_write_bytes", metrics.Counter, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.StreamsWriteSize, err = registry.NewMetric(
		"webtransport_write_size", metrics.Trend, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.StreamsReadCount, err = registry.NewMetric(
		"webtransport_read_count", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.StreamsReadBytes, err = registry.NewMetric(
		"webtransport_read_bytes", metrics.Counter, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.StreamsReadSize, err = registry.NewMetric(
		"webtransport_read_size", metrics.Trend, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.StreamsTotal, err = registry.NewMetric(
		"webtransport_streams_total", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.DatagramsSentCount, err = registry.NewMetric(
		"webtransport_datagrams_sent_count", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.DatagramsSentBytes, err = registry.NewMetric(
		"webtransport_datagrams_sent_bytes", metrics.Counter, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.DatagramsRecvCount, err = registry.NewMetric(
		"webtransport_datagrams_received_count", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.DatagramsRecvBytes, err = registry.NewMetric(
		"webtransport_datagrams_received_bytes", metrics.Counter, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	return wtMetrics, err
}
