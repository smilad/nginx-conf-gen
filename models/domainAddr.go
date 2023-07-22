package models

import (
	"gorm.io/gorm"
	"time"
)

type DomainAddr struct {
	ID              int64           `gorm:"primaryKey,autoIncrement"`
	Name            string          `gorm:"unique"`
	RateLimitConfig RateLimitConfig `gorm:"serializer:json"`
	CacheZone       CacheZone       `gorm:"foreignKey:CacheZoneId"`
	CacheZoneId     int64
	CacheKey        string
	Address         string
	CreatedAt       time.Time      `gorm:"default:now()"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type RateLimitConfig struct {
	Zone    string
	Burst   int
	Rate    string
	MaxSize string
	Path    string
}
