package confgen

var configTemplate = `
proxy_cache_path {{.CacheZone.Path}} levels=1:2 keys_zone={{.CacheZone.ZoneName}}:{{.CacheZone.MaxSize}} inactive={{.CacheZone.Inactive}};
{{ if .RateLimit }}
limit_req_zone $binary_remote_addr zone=addr:{{.RateLimit.MaxSize}} rate={{ .RateLimit.Rate }};
{{ end }}
server {
    listen 80;
    server_name {{ .Domain }};
	{{if .CacheKey}}
	proxy_cache_key $host$request_uri{{.CacheKey.Key}};
	{{end}}

    location / {
        proxy_pass {{ .Addr }};
        proxy_cache {{ .CacheZone.ZoneName }};
        add_header X-Proxy-Cache $upstream_cache_status;
		{{ if .RateLimit }}
        limit_req zone={{.RateLimit.Zone}} burst={{.RateLimit.Burst}};
		{{ end }}
    }
}
`
