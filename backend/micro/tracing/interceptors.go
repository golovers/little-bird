package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
)

// UnaryClientInterceptor creates an interceptor always referring to the active global tracer at time of each call
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(global))
}

// StreamClientInterceptor creates an interceptor always referring to the active global tracer at time of each call
func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(global))
}

// UnaryServerInterceptor creates an interceptor always referring to the active global tracer at time of each call
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(global))
}

// StreamServerInterceptor creates an interceptor always referring to the active global tracer at time of each call
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(global))
}

var global = &globalDelegateTracer{}
// This type introduced in order to lazily delegate to what is the global tracer at time of each call, not at setup
type globalDelegateTracer struct{}

func (*globalDelegateTracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	return opentracing.GlobalTracer().StartSpan(operationName, opts...)
}

func (*globalDelegateTracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	return opentracing.GlobalTracer().Inject(sm, format, carrier)
}

func (*globalDelegateTracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	return opentracing.GlobalTracer().Extract(format, carrier)
}
