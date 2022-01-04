package request

import (
	"goodjobs/business/reviews"
)

type ReviewRequest struct {
	User_ID     uint	`json:"user_id"`
	Building_ID uint	`json:"building_id"`
	Rating      int		`json:"rating"`
	Description string	`json:"description"`
}

func (Review *ReviewRequest) ToDomain() *reviews.Domain{
	return &reviews.Domain{
		User_ID			: Review.User_ID,
		Building_ID		: Review.Building_ID,
		Rating			: Review.Rating,
		Description		: Review.Description,
	}
}