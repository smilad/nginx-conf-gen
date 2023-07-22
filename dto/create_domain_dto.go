package dto

import (
	"fmt"
	"nginx/models"
)

type CreateDomainRequest struct {
	Name      string    `json:"name"   validate:"required,hostname"`
	Address   string    `json:"address" validate:"required"`
	ZoneId    int64     `json:"zoneId" validate:"required"`
	CacheKey  string    `json:"cacheKey" `
	RateLimit rateLimit `json:"rateLimit"`
}

func (r CreateDomainRequest) MapToModel() models.DomainAddr {
	return models.DomainAddr{
		Name:    r.Name,
		Address: r.Address,
		RateLimitConfig: models.RateLimitConfig{
			Zone:    r.RateLimit.Zone,
			Burst:   r.RateLimit.Burst,
			Rate:    fmt.Sprintf("%dr/s", r.RateLimit.RatePerSecond),
			Path:    r.RateLimit.path,
			MaxSize: r.RateLimit.MaxSize,
		},
		CacheZoneId: r.ZoneId,
		CacheKey:    r.CacheKey,
	}
}

type CreateDomainResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
