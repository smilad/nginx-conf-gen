package models

import "time"

type CacheZone struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Path      string
	MaxSize   int
	Inactive  string
	CreatedAt time.Time `gorm:"default:now()"`
}
