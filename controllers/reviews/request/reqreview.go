package request

import (
	"fmt"
	"goodjobs/business/reviews"
)

type ReviewRequest struct {
	User_ID     uint
	Building_ID uint
	Rating      int
	Description string
}

func (Review *ReviewRequest) ToDomain() *reviews.Domain{
	fmt.Println("ini apa", &reviews.Domain{
		User_ID			: Review.User_ID,
		Building_ID		: Review.Building_ID,
		Rating			: Review.Rating,
		Description		: Review.Description,
	})
	fmt.Println("TODOMAIN------->-",Review)
	return &reviews.Domain{
		User_ID			: Review.User_ID,
		Building_ID		: Review.Building_ID,
		Rating			: Review.Rating,
		Description		: Review.Description,
	}
}