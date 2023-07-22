package repository

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"nginx/models"
	"nginx/pkg/confgen"
	"nginx/pkg/filewriter"
	"os"
	"strconv"
)

type ConfigGeneratorRepo struct{}

func NewConfigGeneratorRepo() ConfigGeneratorRepo {
	return ConfigGeneratorRepo{}
}

func (r ConfigGeneratorRepo) GenerateConfig(ctx context.Context, domain models.DomainAddr) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "GenerateConfig[repo.file]")
	defer span.Finish()

	conf := confgen.New(domain.Name, domain.Address)
	conf.SetCacheZone(domain.CacheZone.Name, strconv.Itoa(domain.CacheZone.MaxSize), domain.CacheZone.Inactive, domain.CacheZone.Path)
	conf.SetCacheKey(domain.CacheKey)
	conf.SetRateLimit(domain.RateLimitConfig.Zone, domain.RateLimitConfig.Rate, domain.RateLimitConfig.MaxSize, domain.RateLimitConfig.Burst)

	strConf := conf.Build()
	wd, _ := os.Getwd()
	path := wd + "/generated_configs/" + domain.Name + ".conf"
	file := filewriter.NewFile(path, os.O_CREATE|os.O_RDWR)

	_, err := file.Write([]byte(strConf))
	if err != nil {
		return err
	}

	return nil
}

func (r ConfigGeneratorRepo) DeleteConfig(ctx context.Context, domain string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "GenerateConfig[repo.file]")
	defer span.Finish()

	wd, _ := os.Getwd()
	path := wd + "/generated_configs/" + domain + ".conf"

	return os.Remove(path)
}
