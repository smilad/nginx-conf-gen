package dto

import (
	"nginx/models"
	"time"
)

type GetDomainListResponse struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (d Domain) MapFromEntity(addr models.DomainAddr) Domain {
	return Domain{
		ID:        addr.ID,
		Name:      addr.Name,
		CreatedAt: addr.CreatedAt,
		UpdatedAt: addr.UpdatedAt,
	}
}
