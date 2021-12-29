package request

import (
	"goodjobs/business/reviews"
)

type ReviewRequest struct {
	User_ID     uint
	Building_ID uint
	Rating      int
	Description string
}

func (Review *ReviewRequest) ToDomain() *reviews.Domain{
	return &reviews.Domain{
		User_ID			: Review.User_ID,
		Building_ID		: Review.Building_ID,
		Rating			: Review.Rating,
		Description		: Review.Description,
	}
}