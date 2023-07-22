package dto

import (
	"fmt"
	"nginx/models"
)

type AddCacheZoneRequest struct {
	ZoneName       string `json:"zoneName" validate:"required"`
	MaxSizeMB      int    `json:"maxSizeMB" validate:"required"`
	InactiveSecond int    `json:"inactiveSecond" validate:"required"`
	Path           string `json:"path" validate:"required"`
}

func (m AddCacheZoneRequest) MapToModel() models.CacheZone {
	return models.CacheZone{
		Name:     m.ZoneName,
		MaxSize:  m.MaxSizeMB,
		Inactive: fmt.Sprintf("%ds", m.InactiveSecond),
		Path:     m.Path,
	}
}
