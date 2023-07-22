package admin

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"nginx/models"
	"nginx/pkg/mocks"
	"testing"
	"time"
)

func TestDomainService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()

	domainAddr := models.DomainAddr{
		ID:   1,
		Name: "example.com",
		RateLimitConfig: models.RateLimitConfig{
			Zone:    "example_zone",
			Burst:   10,
			Rate:    "10r/s",
			MaxSize: "100m",
			Path:    "/var/log/nginx",
		},
		CacheZone: models.CacheZone{
			ID:       1,
			Name:     "example_zone",
			Path:     "/var/cache/nginx",
			MaxSize:  100,
			Inactive: "60s",
		},
		CacheZoneId: 1,
		CacheKey:    "example_key",
		Address:     "192.168.1.1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	zone := models.CacheZone{
		ID:       1,
		Name:     "example_zone",
		Path:     "/var/cache/nginx",
		MaxSize:  100,
		Inactive: "60s",
	}

	mockZoneRepo.EXPECT().Get(gomock.Any(), domainAddr.CacheZoneId).Return(zone, nil)
	mockDomainRepo.EXPECT().GetByName(gomock.Any(), domainAddr.Name).Return(false, nil)
	mockDomainRepo.EXPECT().Save(gomock.Any(), domainAddr).Return(domainAddr, nil)
	mockFileRepo.EXPECT().GenerateConfig(gomock.Any(), domainAddr).Return(nil)
	result, err := service.Create(ctx, domainAddr)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDomainService_Create_DuplicateName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()
	zone := models.CacheZone{
		ID:       1,
		Name:     "example_zone",
		Path:     "/var/cache/nginx",
		MaxSize:  100,
		Inactive: "60s",
	}

	// Test data
	domainAddr := models.DomainAddr{
		Name:            "test.com",
		RateLimitConfig: models.RateLimitConfig{},
		CacheZone:       models.CacheZone{},
		CacheZoneId:     1,
		CacheKey:        "",
		Address:         "",
	}

	mockZoneRepo.EXPECT().Get(gomock.Any(), domainAddr.CacheZoneId).Return(zone, nil)
	mockDomainRepo.EXPECT().GetByName(gomock.Any(), domainAddr.Name).Return(true, nil)

	// Call the method to be tested
	_, err := service.Create(ctx, domainAddr)

	// Assert the error
	assert.Error(t, err)
	assert.EqualError(t, err, "record exist")
}

func TestDomainService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()
	domainID := int64(1)
	domainAddr := models.DomainAddr{
		ID:   1,
		Name: "example.com",
		RateLimitConfig: models.RateLimitConfig{
			Zone:    "example_zone",
			Burst:   10,
			Rate:    "10r/s",
			MaxSize: "100m",
			Path:    "/var/log/nginx",
		},
		CacheZone: models.CacheZone{
			ID:       1,
			Name:     "example_zone",
			Path:     "/var/cache/nginx",
			MaxSize:  100,
			Inactive: "60s",
		},
		CacheZoneId: 1,
		CacheKey:    "example_key",
		Address:     "192.168.1.1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockDomainRepo.EXPECT().Delete(gomock.Any(), domainID).Return(domainAddr, nil)

	mockFileRepo.EXPECT().DeleteConfig(gomock.Any(), domainAddr.Name).Return(nil)

	// Call the method to be tested
	err := service.Delete(ctx, domainID)

	// Assert the result
	assert.NoError(t, err)
}

func TestDomainService_Delete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()

	// Test data
	domainID := int64(1)

	// Mock domainRepo.Delete call to return an error
	mockDomainRepo.EXPECT().Delete(gomock.Any(), domainID).Return(models.DomainAddr{}, errors.New("delete error"))

	// Call the method to be tested
	err := service.Delete(ctx, domainID)

	// Assert the error
	assert.Error(t, err)
	assert.EqualError(t, err, "delete error")
}

func TestDomainService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()

	// Test data
	expectedDomains := []models.DomainAddr{
		{
			Name:            "example.com",
			RateLimitConfig: models.RateLimitConfig{},
			CacheZone:       models.CacheZone{},
			CacheZoneId:     1,
			CacheKey:        "",
			Address:         "",
			// Set other necessary fields for domain address
		},
	}

	// Mock domainRepo.GetAll call
	mockDomainRepo.EXPECT().GetAll(gomock.Any()).Return(expectedDomains, nil)

	// Call the method to be tested
	domains, err := service.GetAll(ctx)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedDomains, domains)
}

func TestDomainService_CreateZone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()

	zone := models.CacheZone{
		ID:       1,
		Name:     "example_zone",
		Path:     "/var/cache/nginx",
		MaxSize:  100,
		Inactive: "60s",
	}

	// Mock zoneRepo.Save call
	mockZoneRepo.EXPECT().Save(gomock.Any(), zone).Return(zone, nil)

	// Call the method to be tested
	result, err := service.CreateZone(ctx, zone)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDomainService_GetAllZone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()

	// Test data
	expectedZones := []models.CacheZone{
		{
			ID:       1,
			Name:     "example_zone",
			Path:     "/var/cache/nginx",
			MaxSize:  100,
			Inactive: "60s",
		},
	}

	// Mock zoneRepo.GetAll call
	mockZoneRepo.EXPECT().GetAll(gomock.Any()).Return(expectedZones, nil)

	// Call the method to be tested
	zones, err := service.GetAllZone(ctx)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedZones, zones)
}
