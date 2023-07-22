package repository_contract

import (
	"context"
	"nginx/models"
)

//go:generate mockgen -destination=../pkg/mocks/domain_repository.go --build_flags=--mod=mod -package=mocks . IDomainRepository
type IDomainRepository interface {
	Save(ctx context.Context, e models.DomainAddr) (models.DomainAddr, error)
	Delete(ctx context.Context, id int64) (models.DomainAddr, error)
	GetAll(ctx context.Context) ([]models.DomainAddr, error)
	GetByName(ctx context.Context, name string) (bool, error)
}
