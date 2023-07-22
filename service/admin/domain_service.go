package admin

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"log"
	"nginx/models"
	repo "nginx/repository_contract"
)

type DomainService struct {
	domainRepo repo.IDomainRepository
	zoneRepo   repo.IZoneRepository
	fileRepo   repo.IGenerator
}

func NewDomainService(dr repo.IDomainRepository, zr repo.IZoneRepository, fr repo.IGenerator) DomainService {
	return DomainService{domainRepo: dr, zoneRepo: zr, fileRepo: fr}
}

func (s DomainService) Create(ctx context.Context, e models.DomainAddr) (models.DomainAddr, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Create[service.admin]")
	defer span.Finish()
	zone, err := s.zoneRepo.Get(spannedCtx, e.CacheZoneId)
	if err != nil {
		log.Println(err)
		return models.DomainAddr{}, errors.New("zone not found")
	}

	if find, _ := s.domainRepo.GetByName(spannedCtx, e.Name); find {
		return models.DomainAddr{}, errors.New("record exist")
	}

	res, err := s.domainRepo.Save(spannedCtx, e)
	if err != nil {
		return models.DomainAddr{}, err
	}
	res.CacheZone = zone

	err = s.fileRepo.GenerateConfig(spannedCtx, res)
	if err != nil {
		span.LogKV("error", err)
		return models.DomainAddr{}, errors.New("there is problem in creating config file")
	}
	return res, nil
}

func (s DomainService) Delete(ctx context.Context, domainId int64) error {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Delete[service.admin]")
	defer span.Finish()

	m, err := s.domainRepo.Delete(spannedCtx, domainId)
	if err != nil {
		return err
	}
	err = s.fileRepo.DeleteConfig(spannedCtx, m.Name)
	if err != nil {
		log.Println(err)
		return errors.New("file deleting is not successfully")
	}
	return nil
}

func (s DomainService) GetAll(ctx context.Context) ([]models.DomainAddr, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "Delete[service.admin]")
	defer span.Finish()

	return s.domainRepo.GetAll(spannedCtx)
}

func (s DomainService) CreateZone(ctx context.Context, zone models.CacheZone) (models.CacheZone, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "CreateZone[service.admin]")
	defer span.Finish()

	return s.zoneRepo.Save(spannedCtx, zone)
}

func (s DomainService) GetAllZone(ctx context.Context) ([]models.CacheZone, error) {
	span, spannedCtx := opentracing.StartSpanFromContext(ctx, "GetAllZone[service.admin]")
	defer span.Finish()

	return s.zoneRepo.GetAll(spannedCtx)
}
