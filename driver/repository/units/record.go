package units

import (
	"goodjobs/business/units"
	"goodjobs/driver/repository/buildings"
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	Id          uint				`gorm:"primaryKey"`
	Name        string				`gorm:"unique"`
	Building_ID uint
	Building    buildings.Building	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description string
	Price       uint
	UnitSize    float64
	Img         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt 		`gorm:"index"`
}

func (unit *Unit) ToDomain() units.Domain {
	res := units.Domain{
		Id					:unit.Id,
		Name 				:unit.Name,
		Building_ID			:unit.Building_ID,
		Building			:unit.Building.ToDomain(),
		Description			:unit.Description,
		Price				:unit.Price,
		UnitSize			:unit.UnitSize,
		Img					:unit.Img,
		CreatedAt			:unit.CreatedAt,
		UpdatedAt			:unit.UpdatedAt,
	}
	return res
}

func FromDomain(domain units.Domain) Unit {
	return Unit{
		Id					:domain.Id,
		Name				:domain.Name,
		Building_ID			:domain.Building_ID,
		Building			:buildings.FromDomain(domain.Building),
		Description			:domain.Description ,
		Price				:domain.Price,
		UnitSize			:domain.UnitSize,
		Img					:domain.Img,
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}

}

func ToDomainArray(data []Unit) []units.Domain{
	res := []units.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}
