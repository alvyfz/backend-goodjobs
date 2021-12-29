package response

import (
	"goodjobs/business/reviews"
	respBuilding "goodjobs/controllers/buildings/response"
	respUser "goodjobs/controllers/users/response"
	"time"

	"gorm.io/gorm"
)

type ReviewResponse struct {
	Id          uint `gorm:"primaryKey"`
	User_ID     uint `json:"user_id"`
	User        respUser.UserResponse	`json:"user"`
	Building_ID uint	`json:"building_id"`
	Building	respBuilding.BuildingResponse	`json:"building"`
	Rating      int				`json:"rating"`
	Description string				`json:"description"`
	CreatedAt   time.Time			
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func FromDomainReview(domain reviews.Domain) ReviewResponse {
	return ReviewResponse{
		Id			:domain.Id,
		User_ID		:domain.User_ID,
		User		:respUser.FromDomain(domain.User),
		Building_ID	:domain.Building_ID,
		Building	:respBuilding.FromDomainBuilding(domain.Building),
		Rating		:domain.Rating,
		Description: domain.Description,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}
}

func FromDomainReviewArray(data []reviews.Domain) []ReviewResponse{
	var res []ReviewResponse
	for _, val := range data{
		res = append(res, FromDomainReview(val))
	}
	return res
}