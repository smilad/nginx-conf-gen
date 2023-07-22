package config

import "time"

// Service details
type Service struct {
	Name      string `yaml:"name" required:"true"`
	ID        uint32 `yaml:"id" required:"true"`
	BaseURL   string `yaml:"baseURL"`
	Server    Server `yaml:"service.server"`
	MetricUrl string `yaml:"metricUrl"`
	Debug     bool   `yaml:"debug"`
}

type Server struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
}
