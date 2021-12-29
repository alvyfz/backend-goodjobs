package units

import (
	"goodjobs/business/units"
	"goodjobs/controllers"
	"goodjobs/controllers/units/request"
	"goodjobs/controllers/units/response"
	"goodjobs/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UnitController struct {
	unitUseCase units.UnitUsecaseInterface
}

func NewUnitController(UnitUseCase units.UnitUsecaseInterface) *UnitController {
	return &UnitController{
		unitUseCase: UnitUseCase,
	}
}

func (unitController *UnitController) Add(c echo.Context) error {
	req := request.UnitRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data units.Domain
	data, err = unitController.unitUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainUnit(data))

}

func (unitController *UnitController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	unit, err := unitController.unitUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainUnitArray(unit))

}

func (unitController *UnitController) GetByID(c echo.Context) error{
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := unitController.unitUseCase.GetByID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainUnit(data))
}

func (unitController *UnitController) GetByBuildingID(c echo.Context) error{
	req := c.Request().Context()
	buildingid := c.Param("buildingid")
	Convint, errConvint := strconv.Atoi(buildingid)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := unitController.unitUseCase.GetByBuildingID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainUnitArray(data))
}

func (unitController *UnitController) Edit (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.UnitRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := unitController.unitUseCase.Edit(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainUnit(data))

}

func (unitController *UnitController) Delete(c echo.Context) error {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = unitController.unitUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}