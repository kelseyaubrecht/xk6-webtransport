package wt

import (
	"errors"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type WTMetrics struct {
	WriteCount *metrics.Metric
	WriteBytes *metrics.Metric
	WriteSize  *metrics.Metric

	ReadCount *metrics.Metric
	ReadBytes *metrics.Metric
	ReadSize  *metrics.Metric
}

func registerMetrics(vu modules.VU) (WTMetrics, error) {
	var err error
	registry := vu.InitEnv().Registry
	wtMetrics := WTMetrics{}

	if wtMetrics.WriteCount, err = registry.NewMetric(
		"webtransport_write_count", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.WriteBytes, err = registry.NewMetric(
		"webtransport_write_bytes", metrics.Counter, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.WriteSize, err = registry.NewMetric(
		"webtransport_write_size", metrics.Trend, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.ReadCount, err = registry.NewMetric(
		"webtransport_read_count", metrics.Counter); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.ReadBytes, err = registry.NewMetric(
		"webtransport_read_bytes", metrics.Counter, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	if wtMetrics.ReadSize, err = registry.NewMetric(
		"webtransport_read_size", metrics.Trend, metrics.Data); err != nil {
		return wtMetrics, errors.Unwrap(err)
	}

	return wtMetrics, err
}
