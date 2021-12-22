package complexes

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Img		  string
}

type ComplexUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context) error
}

type ComplexRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context) error
}