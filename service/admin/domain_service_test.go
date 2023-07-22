package admin

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"nginx/models"
	"nginx/pkg/mocks"
)

func TestDomainService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDomainRepo := mocks.NewMockIDomainRepository(ctrl)
	mockZoneRepo := mocks.NewMockIZoneRepository(ctrl)
	mockFileRepo := mocks.NewMockIGenerator(ctrl)

	service := NewDomainService(mockDomainRepo, mockZoneRepo, mockFileRepo)

	ctx := context.Background()

	// Test data
	domainAddr := models.DomainAddr{
		Name: "test.com",
		RateLimitConfig: models.RateLimitConfig{
			Zone:    "1",
			Burst:   54,
			Rate:    "",
			MaxSize: "",
			Path:    "",
		},
		CacheZone: models.CacheZone{
			ID:        0,
			Name:      "",
			Path:      "",
			MaxSize:   0,
			Inactive:  "",
			CreatedAt: time.Time{},
		},
		CacheZoneId: 0,
		CacheKey:    "",
		Address:     "",
	}

	zone := models.CacheZone{
		ID:      1,
		Name:    "example_zone",
		MaxSize: 100,
	}

	// Mock zoneRepo.Get call
	mockZoneRepo.EXPECT().Get(gomock.Any(), domainAddr.CacheZoneId).Return(zone, nil)

	// Mock domainRepo.GetByName call
	mockDomainRepo.EXPECT().GetByName(gomock.Any(), domainAddr.Name).Return(false, nil)

	// Mock domainRepo.Save call
	mockDomainRepo.EXPECT().Save(gomock.Any(), domainAddr).Return(domainAddr, nil)

	// Mock fileRepo.GenerateConfig call
	mockFileRepo.EXPECT().GenerateConfig(gomock.Any(), domainAddr).Return(nil)

	// Call the method to be tested
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

	// Test data
	domainAddr := models.DomainAddr{
		Name:            "test.com",
		RateLimitConfig: models.RateLimitConfig{},
		CacheZone:       models.CacheZone{},
		CacheZoneId:     0,
		CacheKey:        "",
		Address:         "",
	}

	// Mock domainRepo.GetByName call to return true, indicating the name already exists
	mockDomainRepo.EXPECT().GetByName(gomock.Any(), domainAddr.Name).Return(true, nil)

	// Call the method to be tested
	_, err := service.Create(ctx, domainAddr)

	// Assert the error
	assert.Error(t, err)
	assert.EqualError(t, err, "record exist")
}
