package complexes

import (
	"goodjobs/business/complexes"
	"goodjobs/controllers"
	"goodjobs/controllers/complexes/request"
	"goodjobs/controllers/complexes/response"
	"goodjobs/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ComplexController struct {
	complexUseCase complexes.ComplexUsecaseInterface
}

func NewComplexController(ComplexUseCase complexes.ComplexUsecaseInterface) *ComplexController{
	return &ComplexController{
		complexUseCase: ComplexUseCase,
	}
}

func (complexController *ComplexController) Add(c echo.Context) error {
	req := request.ComplexRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data complexes.Domain
	data, err = complexController.complexUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainComplex(data))

}

func (complexController *ComplexController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	complex, err := complexController.complexUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAll(complex))

}

func (complexController *ComplexController) GetByID(c echo.Context) error{
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := complexController.complexUseCase.GetByID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainComplex(data))
}

func (complexController *ComplexController) Edit (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.ComplexRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := complexController.complexUseCase.Edit(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainComplex(data))

}

func (complexController *ComplexController) Delete(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = complexController.complexUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}