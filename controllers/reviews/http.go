package reviews

import (
	"fmt"
	"goodjobs/business/reviews"
	"goodjobs/controllers"
	"goodjobs/controllers/reviews/request"
	"goodjobs/controllers/reviews/response"
	"goodjobs/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	reviewUseCase reviews.ReviewUsecaseInterface
}

func NewReviewController(ReviewUseCase reviews.ReviewUsecaseInterface) *ReviewController {
	return &ReviewController{
		reviewUseCase: ReviewUseCase,
	}
}

func (reviewController *ReviewController) Add(c echo.Context) error {
	req := request.ReviewRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data reviews.Domain
	fmt.Println(req.ToDomain(), "ctx", ctx)
	data, err = reviewController.reviewUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainReview(data))

}

func (reviewController *ReviewController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	Review, err := reviewController.reviewUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAll(Review))

}

func (reviewController *ReviewController) GetByBuildingID(c echo.Context) error{
	req := c.Request().Context()
	buildingid := c.Param("buildingid")
	Convint, errConvint := strconv.Atoi(buildingid)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := reviewController.reviewUseCase.GetByBuildingID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainReview(data))
}

func (reviewController *ReviewController) Edit (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.ReviewRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := reviewController.reviewUseCase.Edit(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainReview(data))

}


func (reviewController *ReviewController) Delete(c echo.Context) error {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = reviewController.reviewUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}
