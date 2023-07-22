package repository

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"nginx/models"
)

type ZoneRepository struct {
	db *gorm.DB
}

func NewZoneRepository(db *gorm.DB) ZoneRepository {
	return ZoneRepository{
		db: db,
	}
}

func (z ZoneRepository) Save(ctx context.Context, m models.CacheZone) (models.CacheZone, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Save[repository.zone]")
	defer span.Finish()
	if err := z.db.WithContext(spannedCtx).Save(&m).Error; err != nil {
		span.SetTag("error", true)
		span.LogKV("error-log", err)
		return models.CacheZone{}, err
	}
	return m, nil
}

func (z ZoneRepository) GetAll(ctx context.Context) ([]models.CacheZone, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Save[repository.zone]")
	defer span.Finish()
	var zones []models.CacheZone
	if err := z.db.WithContext(spannedCtx).Find(&zones).Error; err != nil {
		span.SetTag("error", true)
		span.LogKV("error-log", err)
		return nil, err
	}
	return zones, nil
}

func (z ZoneRepository) Get(ctx context.Context, id int64) (models.CacheZone, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Save[repository.zone]")
	defer span.Finish()
	var zones models.CacheZone
	if err := z.db.WithContext(spannedCtx).First(&zones, id).Error; err != nil {
		span.SetTag("error", true)
		span.LogKV("error-log", err)
		return models.CacheZone{}, err
	}
	return zones, nil
}
