package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"io"
	"log"
	"nginx/config"
)

// Connect method
func InitTracer(confs *config.Config) (io.Closer, error) {
	// Initialize tracer with a logger and a metrics factory
	var err error

	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via  configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: confs.Service.Name,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           confs.Jaeger.LogSpans,
			LocalAgentHostPort: confs.Jaeger.HostPort,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
		jaegercfg.ZipkinSharedRPCSpan(true),
	)

	if err != nil {
		log.Fatalf("error in initial config")
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, err
}
