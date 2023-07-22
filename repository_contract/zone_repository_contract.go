package repository_contract

import (
	"context"
	"nginx/models"
)

//go:generate mockgen -destination=../pkg/mocks/zone_repository.go --build_flags=--mod=mod -package=mocks . IZoneRepository
type IZoneRepository interface {
	Save(ctx context.Context, m models.CacheZone) (models.CacheZone, error)
	GetAll(ctx context.Context) ([]models.CacheZone, error)
	Get(ctx context.Context, id int64) (models.CacheZone, error)
}
