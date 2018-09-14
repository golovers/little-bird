package stackdriver

import (
	"context"
	"time"
)

type metricKey struct{}

// Metrics holds per request metrics.
type Metrics struct {
	// Code is the first http response code passed to WriteHeader.
	Code int
	// Duration of the handler.
	Duration time.Duration
	// Written is the number of bytes written to the response writer.
	Written int64
}

// NewContext creates a new context.Context with the metrics attached.
func NewContext(ctx context.Context, m Metrics) context.Context {
	return context.WithValue(ctx, metricKey{}, m)
}

// FromContext extracts the Metrics from the context if available.
func FromContext(ctx context.Context) (m Metrics, ok bool) {
	m, ok = ctx.Value(metricKey{}).(Metrics)
	return
}
