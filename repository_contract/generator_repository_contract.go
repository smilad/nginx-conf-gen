package repository_contract

import (
	"context"
	"nginx/models"
)

//go:generate mockgen -destination=../pkg/mocks/generator_repository.go --build_flags=--mod=mod -package=mocks . IGenerator
type IGenerator interface {
	GenerateConfig(ctx context.Context, domain models.DomainAddr) error
	DeleteConfig(ctx context.Context, domain string) error
}
