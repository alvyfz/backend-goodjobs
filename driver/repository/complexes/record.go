package complexes

import (
	"goodjobs/business/complexes"
	"time"

	"gorm.io/gorm"
)

type Complex struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Address	  string
	Img       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (complex *Complex) ToDomain() complexes.Domain{
	res := complexes.Domain{
		Id: complex.Id,
		Name: complex.Name,
		Address: complex.Address,
		Img: complex.Img,
		CreatedAt 	:complex.CreatedAt,
		UpdatedAt 	:complex.UpdatedAt,
	}
	return res
}

func FromDomain(domain complexes.Domain) Complex {
	return Complex{
		Id: domain.Id,
		Name: domain.Name,
		Address: domain.Address,
		Img: domain.Img,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}
}

func GetAll(data []Complex) []complexes.Domain{
	res := []complexes.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}