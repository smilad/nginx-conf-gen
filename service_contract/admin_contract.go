package service_contract

import (
	"context"
	"nginx/models"
)

//go:generate mockgen -destination=../pkg/mocks/admin_service.go -package=mocks . IDomainService
type IDomainService interface {
	Create(ctx context.Context, e models.DomainAddr) (models.DomainAddr, error)
	Delete(ctx context.Context, domainId int64) error
	GetAll(ctx context.Context) ([]models.DomainAddr, error)
	CreateZone(ctx context.Context, zone models.CacheZone) (models.CacheZone, error)
	GetAllZone(ctx context.Context) ([]models.CacheZone, error)
}
