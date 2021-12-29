package request

import "goodjobs/business/units"

type UnitRequest struct {
	Name        string `json:"name"`
	Building_ID uint   `json:"building_id"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	UnitSize    float64   `json:"unitsize"`
	Img         string `json:"img"`
}

func (Unit *UnitRequest) ToDomain() *units.Domain {
	return &units.Domain{
		Name:        Unit.Name,
		Building_ID: Unit.Building_ID,
		Description: Unit.Description,
		Price:       Unit.Price,
		UnitSize:    Unit.UnitSize,
		Img:         Unit.Img,
	}
}