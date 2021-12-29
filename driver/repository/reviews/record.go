package reviews

import (
	"goodjobs/business/reviews"
	"goodjobs/driver/repository/buildings"
	"goodjobs/driver/repository/users"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	Id        	uint   `gorm:"primaryKey"`
	User_ID     uint
	User		users.User	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Building_ID	uint
	Building	buildings.Building	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Rating      int
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (review *Review) ToDomain() reviews.Domain{
	res := reviews.Domain{
		Id			:review.Id,
		User_ID		:review.User_ID,
		User		:review.User.ToDomain(),
		Building_ID :review.Building_ID,
		Building	:review.Building.ToDomain(),
		Rating		:review.Rating,
		Description	:review.Description,
		CreatedAt 	:review.CreatedAt,
		UpdatedAt 	:review.UpdatedAt,
	}
	return res
}

func FromDomain(domain reviews.Domain) Review {
	return Review{
		Id			:domain.Id,
		User_ID		:domain.User_ID,
		User		:users.FromDomain(domain.User),
		Building_ID	:domain.Building_ID,
		Building	:buildings.FromDomain(domain.Building),
		Rating		:domain.Rating,
		Description: domain.Description,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}
}

func GetAll(data []Review) []reviews.Domain{
	res := []reviews.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}