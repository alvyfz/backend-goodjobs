package buildings

import (
	"context"
	"goodjobs/business/complexes"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Complex_ID	uint
	Complex		complexes.Domain
	Name        string
	Description string
	Size 		uint
	Floor 		int
	OfficeHours string
	Address	 	string
	Toilet 		uint
	Img 		string
	Latitude 	float64
	Longitude 	float64
	PriceStart	uint
}

type BuildingUsecaseInterface interface{
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	GetByComplexID(complexid uint, ctx context.Context) ([]Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

type BuildingRepoInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	GetByComplexID(complexid uint, ctx context.Context) ([]Domain, error)
	Edit(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}