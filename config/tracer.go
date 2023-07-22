package config

// Jaeger tracer
type tracer struct {
	HostPort string `yaml:"hostPort" required:"true"`
	LogSpans bool   `yaml:"logSpans"`
}
