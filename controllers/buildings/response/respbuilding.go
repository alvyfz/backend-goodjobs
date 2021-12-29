package response

import (
	"goodjobs/business/buildings"
	respComp "goodjobs/controllers/complexes/response"
	"time"

	"gorm.io/gorm"
)

type BuildingResponse struct {
	Id          uint 	`gorm:"primaryKey"`
	Complex_ID  uint	`json:"complex_id"`
	Complex     respComp.ComplexResponse	`json:"complex"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Size        float64    `json:"size"`
	Floor       int     `json:"floor"`
	OfficeHours string  `json:"officehours"`
	Address     string  `json:"address"`
	Toilet      uint    `json:"toilet"`
	Img         string  `json:"img"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	PriceStart	uint	`json:"pricestart"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromDomainBuilding(domain buildings.Domain) BuildingResponse {
	return BuildingResponse{
		Id					:domain.Id,
		Complex_ID			:domain.Complex_ID,
		Complex				:respComp.FromDomainComplex(domain.Complex),
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

func FromDomainBuildingArray(data []buildings.Domain) []BuildingResponse{
	var res []BuildingResponse
	for _, val := range data{
		res = append(res, FromDomainBuilding(val))
	}
	return res
}