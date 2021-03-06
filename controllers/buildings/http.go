package buildings

import (
	"goodjobs/business/buildings"
	"goodjobs/controllers"
	"goodjobs/controllers/buildings/request"
	"goodjobs/controllers/buildings/response"
	"goodjobs/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BuildingController struct {
	buildingUseCase buildings.BuildingUsecaseInterface
}

func NewBuildingController(BuildingUseCase buildings.BuildingUsecaseInterface) *BuildingController {
	return &BuildingController{
		buildingUseCase: BuildingUseCase,
	}
}

func (buildingController *BuildingController) Add(c echo.Context) error {
	req := request.BuildingRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data buildings.Domain
	data, err = buildingController.buildingUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuilding(data))

}

func (buildingController *BuildingController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	Building, err := buildingController.buildingUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuildingArray(Building))

}

func (buildingController *BuildingController) GetOrderByPriceAsc(c echo.Context) error {
	req := c.Request().Context()
	Building, err := buildingController.buildingUseCase.GetOrderByPriceAsc(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuildingArray(Building))

}

func (buildingController *BuildingController) GetOrderByPriceDesc(c echo.Context) error {
	req := c.Request().Context()
	Building, err := buildingController.buildingUseCase.GetOrderByPriceDesc(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuildingArray(Building))

}

func (buildingController *BuildingController) GetByID(c echo.Context) error{
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := buildingController.buildingUseCase.GetByID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuilding(data))
}

func (buildingController *BuildingController) GetByComplexID(c echo.Context) error{
	req := c.Request().Context()
	complexid := c.Param("complexid")
	Convint, errConvint := strconv.Atoi(complexid)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := buildingController.buildingUseCase.GetByComplexID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuildingArray(data))
}

func (buildingController *BuildingController) Edit (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.BuildingRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := buildingController.buildingUseCase.Edit(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainBuilding(data))

}


func (buildingController *BuildingController) Delete(c echo.Context) error {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = buildingController.buildingUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}