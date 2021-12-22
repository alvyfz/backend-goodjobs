package units

import (
	"context"
	"goodjobs/business/buildings"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Building_ID uint
	Building	buildings.Domain
	Description string
	Price       uint
	UnitSize    uint
	Img         string
}

type UnitUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context) error
}

type UnitRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context) error
}