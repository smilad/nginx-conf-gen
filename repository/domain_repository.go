package repository

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"nginx/models"
)

type DomainAddrRepository struct {
	db *gorm.DB
}

func NewDomainRepository(db *gorm.DB) DomainAddrRepository {
	return DomainAddrRepository{
		db: db,
	}
}

func (d DomainAddrRepository) Save(ctx context.Context, e models.DomainAddr) (models.DomainAddr, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Save[repository.domain]")
	defer span.Finish()
	if err := d.db.WithContext(spannedCtx).Create(&e).Error; err != nil {
		span.SetTag("error", true)
		span.LogKV("error-saving-domainAdr", err)
		return models.DomainAddr{}, err
	}

	return e, nil
}

func (d DomainAddrRepository) Delete(ctx context.Context, id int64) (m models.DomainAddr, err error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Delete[repository.domain]")
	defer span.Finish()

	result := d.db.WithContext(spannedCtx).First(&m, id).Delete(&m, id)
	if result.Error != nil {
		span.SetTag("error", true)
		span.LogKV("error-delete-domainAdr", result.Error)
		return models.DomainAddr{}, result.Error
	}
	if result.RowsAffected < 1 {
		return models.DomainAddr{}, errors.New("domain not found or deleted")
	}

	return
}

func (d DomainAddrRepository) GetAll(ctx context.Context) ([]models.DomainAddr, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "GetAll[repository.domain]")
	defer span.Finish()
	var result []models.DomainAddr

	if err := d.db.WithContext(spannedCtx).Find(&result).Error; err != nil {
		span.SetTag("error", true)
		span.LogKV("error-Update-domainAdr", err)
		return nil, err
	}

	return result, nil
}

func (d DomainAddrRepository) Get(ctx context.Context, id int64) (models.DomainAddr, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Update[repository.domain]")
	defer span.Finish()
	var result models.DomainAddr
	if err := d.db.Debug().WithContext(spannedCtx).First(&result).Error; err != nil {
		span.SetTag("error", true)
		span.LogKV("error-Update-domainAdr", err)
		return models.DomainAddr{}, nil
	}

	return result, nil
}

func (d DomainAddrRepository) GetByName(ctx context.Context, name string) (bool, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "GetByName[repository.domain]")
	defer span.Finish()
	result := d.db.WithContext(spannedCtx).Debug().Where("name = ?", name).Find(&models.DomainAddr{})
	if result.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
