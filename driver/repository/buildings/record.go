package buildings

import (
	"goodjobs/business/buildings"
	"goodjobs/driver/repository/complexes"
	"time"

	"gorm.io/gorm"
)

type Building struct {
	Id          uint			 `gorm:"primaryKey"`
	Complex_ID  uint
	Complex     complexes.Complex `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Name        string			`gorm:"unique"`
	Description string			
	Size        float64
	Floor       int
	OfficeHours string
	Address     string
	Toilet      uint
	Img         string
	Latitude    float64
	Longitude   float64
	PriceStart	uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (building *Building) ToDomain() buildings.Domain {
	res := buildings.Domain{
		Id:          building.Id,				
		CreatedAt:   building.CreatedAt,
		UpdatedAt:   building.UpdatedAt,
		DeletedAt:   gorm.DeletedAt{},
		Complex_ID:  building.Complex_ID,
		Complex:     building.Complex.ToDomain(),
		Name:        building.Name,
		Description: building.Description,
		Size:        building.Size,
		Floor:       building.Floor,
		OfficeHours: building.OfficeHours,
		Address:     building.Address,
		Toilet:      building.Toilet,
		Img:         building.Img,
		Latitude:    building.Latitude,
		Longitude:   building.Longitude,
		PriceStart:  building.PriceStart,
	}
	return res
}

func FromDomain(domain buildings.Domain) Building {
	return Building{
		Id					:domain.Id,
		Complex_ID			:domain.Complex_ID,
		Complex				:complexes.FromDomain(domain.Complex) ,
		Name				:domain.Name,
		Description			:domain.Description,
		Size				:domain.Size,
		Floor				:domain.Floor,
		OfficeHours			:domain.OfficeHours,
		Address				:domain.Address,
		Toilet				:domain.Toilet,
		Img					:domain.Img,
		Latitude			:domain.Latitude,
		Longitude			:domain.Longitude,
		PriceStart			:domain.PriceStart,
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}
}

func ToDomainArray(data []Building) []buildings.Domain{
	res := []buildings.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}
