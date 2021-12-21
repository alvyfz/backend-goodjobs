package request

import "goodjobs/business/buildings"

type BuildingRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Size        uint    `json:"size"`
	Floor       int     `json:"floor"`
	OfficeHours string  `json:"officehours"`
	Address     string  `json:"address"`
	Toilet      uint    `json:"toilet"`
	Img         string  `json:"img"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func (Building *BuildingRequest) ToDomain() *buildings.Domain {
	return &buildings.Domain{
		Name				:Building.Name,
		Description			:Building.Description,
		Size				:Building.Size,
		Floor				:Building.Floor,
		OfficeHours			:Building.OfficeHours,
		Address				:Building.Address,
		Toilet				:Building.Toilet,
		Img					:Building.Img,
		Latitude			:Building.Latitude,
		Longitude			:Building.Longitude,
	}
}