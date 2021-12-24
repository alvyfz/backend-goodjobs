package response

import (
	"goodjobs/business/complexes"
	"time"

	"gorm.io/gorm"
)

type ComplexResponse struct {
	Id        uint           `json:"id"`
	Name      string         `json:"name"`
	Address	  string 		 `json:"address"`
	Img       string         `json:"img"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func FromDomainComplex(domain complexes.Domain) ComplexResponse {
	return ComplexResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Address:   domain.Address,
		Img:       domain.Img,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func GetAll(domain []complexes.Domain) []ComplexResponse {
	var GetAll []ComplexResponse
	for _, val := range domain {
		GetAll = append(GetAll, FromDomainComplex(val))
	}
	return GetAll
}