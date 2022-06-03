package transport

import (
	"context"
	"fmt"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	metrics "github.com/go-kit/kit/metrics"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	opentracinggo "github.com/opentracing/opentracing-go"
)

type EndpointsSet struct {
	AddEndpoint endpoint.Endpoint
	GetEndpoint endpoint.Endpoint
	// ListEndpoint   endpoint.Endpoint
	// UpdateEndpoint endpoint.Endpoint
}

func InstrumentingEndpoints(endpoints EndpointsSet, tracer opentracinggo.Tracer) EndpointsSet {
	return EndpointsSet{
		AddEndpoint: opentracing.TraceServer(tracer, "Add")(endpoints.AddEndpoint),
		GetEndpoint: opentracing.TraceServer(tracer, "Get")(endpoints.GetEndpoint),
		// ListEndpoint:   opentracing.TraceServer(tracer, "List")(endpoints.ListEndpoint),
		// UpdateEndpoint: opentracing.TraceServer(tracer, "Update")(endpoints.UpdateEndpoint),
	}
}

func LatencyMiddleware(dur metrics.Histogram, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		dur := dur.With("method", methodName)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				dur.With("success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func RequestFrequencyMiddleware(freq metrics.Gauge, methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		freq := freq.With("method", methodName)
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			freq.Add(1)
			response, err := next(ctx, request)
			freq.Add(-1)
			return response, err
		}
	}
}
