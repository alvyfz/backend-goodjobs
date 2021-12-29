package response

import (
	"goodjobs/business/units"
	respBuil "goodjobs/controllers/buildings/response"
	"time"

	"gorm.io/gorm"
)

type UnitResponse struct {
	Id          uint   			`json:"id"`
	Name        string 			`gorm:"unique"`
	Building_ID uint			`json:"building_id"`
	Building	respBuil.BuildingResponse	`json:"building"`
	Description string			`json:"description"`
	Price       uint			`json:"price"`
	UnitSize    uint			`json:"unitsize"`
	Img         string			`json:"img"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt  `json:"deletedAt"`
}

func FromDomainUnit(domain units.Domain) UnitResponse {
	return UnitResponse{
		Id					:domain.Id,
		Name				:domain.Name,
		Building_ID			:domain.Building_ID,
		Building			:respBuil.FromDomainBuilding(domain.Building),
		Description			:domain.Description ,
		Price				:domain.Price,
		UnitSize			:domain.UnitSize,
		Img					:domain.Img,
		CreatedAt			:domain.CreatedAt,
		UpdatedAt			:domain.UpdatedAt,
	}
}

func FromDomainUnitArray(data []units.Domain) []UnitResponse{
	var res []UnitResponse
	for _, val := range data{
		res = append(res, FromDomainUnit(val))
	}
	return res
}