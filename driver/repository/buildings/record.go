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
	Complex     complexes.Complex `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string
	Description string
	Size        uint
	Floor       int
	OfficeHours string
	Address     string
	Toilet      uint
	Img         string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (building *Building) ToDomain() buildings.Domain {
	res := buildings.Domain{
		Id					:building.Id,
		Complex_ID			:building.Complex_ID,
		Complex				:building.ToDomain().Complex,
		Name				:building.Name,
		Description			:building.Description,
		Size				:building.Size,
		Floor				:building.Floor,
		OfficeHours			:building.OfficeHours,
		Address				:building.Address,
		Toilet				:building.Toilet,
		Img					:building.Img,
		Latitude			:building.Latitude,
		Longitude			:building.Longitude,
		CreatedAt			:building.CreatedAt,
		UpdatedAt			:building.UpdatedAt,
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
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}
}

func GetAll(data []Building) []buildings.Domain{
	res := []buildings.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}