package reviews

import (
	"context"
	"goodjobs/business/buildings"
	"goodjobs/business/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id          uint
	User_ID     uint
	User		users.Domain
	Building_ID	uint
	Building	buildings.Domain
	Rating      int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ReviewUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByBuildingID(buildingid uint, ctx context.Context) (Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context) error
}

type ReviewRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByBuildingID(buildingid uint, ctx context.Context) (Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context) error
}