package tracing

import (
	"context"
	"fmt"

	"github.com/opentracing/basictracer-go"
	"github.com/sirupsen/logrus"
	"github.com/opentracing/opentracing-go"
)

// TraceLogger returns a new logrus.Entry, containing trace if configured
func TraceLogger(ctx context.Context) *logrus.Entry {
	return logrus.WithFields(LogFields(ctx))
}

// Log is alias for more verbose call TraceLogger
var Log = TraceLogger

// Get trace information to be provided to logger
func LogFields(ctx context.Context) map[string]interface{} {
	spanContext, ok := opentracing.SpanFromContext(ctx).Context().(basictracer.SpanContext)
	if !ok || traceProjectId == "" {
		return map[string]interface{}{}
	}
	return map[string]interface{} {
		"trace": fmt.Sprintf("projects/%s/traces/%d", traceProjectId, spanContext.TraceID),
	}
}
