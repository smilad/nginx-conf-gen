package confgen

import (
	"bytes"
	"fmt"
	"text/template"
)

type CacheKey struct {
	Key string
}

type CacheZone struct {
	ZoneName string
	MaxSize  string
	Inactive string
	Path     string
}

type RateLimit struct {
	Zone    string
	Burst   int
	Rate    string
	MaxSize string
}

type NginxConfig struct {
	Domain    string
	Addr      string
	CacheKey  *CacheKey
	CacheZone *CacheZone
	RateLimit *RateLimit
}

func New(domain, addr string) NginxConfig {
	return NginxConfig{
		Domain: domain,
		Addr:   addr,
	}
}

func (ng *NginxConfig) SetCacheKey(key string) {
	ng.CacheKey = &CacheKey{Key: key}
}

func (ng *NginxConfig) SetRateLimit(zone, rate, maxsize string, burst int) {
	ng.RateLimit = &RateLimit{
		Zone:    zone,
		Burst:   burst,
		Rate:    rate,
		MaxSize: maxsize,
	}
}

func (ng *NginxConfig) SetCacheZone(zoneName, maxSize, inactive, path string) {
	ng.CacheZone = &CacheZone{
		ZoneName: zoneName,
		MaxSize:  maxSize,
		Inactive: inactive,
		Path:     path,
	}
}

func (ng *NginxConfig) Build() string {
	var c bytes.Buffer
	config := template.Must(template.New("config").Parse(configTemplate))
	fmt.Println(ng)
	_ = config.Execute(&c, ng)

	return c.String()
}
