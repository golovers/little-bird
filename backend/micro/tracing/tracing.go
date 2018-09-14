package tracing

import (
	"context"
	"fmt"

	"cloud.google.com/go/trace/apiv1"
	"github.com/lovoo/gcloud-opentracing"
	"github.com/opentracing/basictracer-go"
	"github.com/sirupsen/logrus"
	"github.com/opentracing/opentracing-go"
)

var traceProjectId = ""

// ConfigureTracing sets up tracing based on supplied config
//
// this sets the global tracer: opentracing.GlobalTracer()
//
// returns false is tracing was not configured
func ConfigureTracing(ctx context.Context, cfg *Config) (bool, error) {
	if cfg.TracingPeriod == 0 {
		logrus.Warnf("tracing will not be configured; tracing period = 0")
		return false, nil
	}

	if cfg.ProjectID == "" {
		logrus.Warnf("tracing will not be configured; tracing project unspecified")
		return false, nil
	}

	logrus.Infof("configuring tracing with google project: %q. Every 1/%d request will be traced",
		cfg.ProjectID, cfg.TracingPeriod)

	client, err := trace.NewClient(ctx)
	if err != nil {
		return false, fmt.Errorf("could not create trace client: %v", err)
	}

	recorder, err := gcloudtracer.NewRecorder(ctx, cfg.ProjectID, client, gcloudtracer.WithLogger(logrus.StandardLogger()))
	if err != nil {
		return false, fmt.Errorf("could not create trace recorder: %v", err)
	}

	opentracing.SetGlobalTracer(basictracer.NewWithOptions(basictracer.Options{
		Recorder:     recorder,
		ShouldSample: shouldSampleFunc(cfg.TracingPeriod),
	}))

	traceProjectId = cfg.ProjectID

	return true, nil
}

func shouldSampleFunc(period uint64) func(traceID uint64) bool {
	if period == 1 {
		return func(uint64) bool { return true }
	}
	return func(traceId uint64) bool { return traceId%period == 0 }
}
