package roles

import (
	"goodjobs/business/roles"
	"goodjobs/controllers"
	"goodjobs/controllers/roles/request"
	"goodjobs/controllers/roles/response"
	"goodjobs/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoleController struct {
	roleUseCase roles.RoleUsecaseInterface
}

func NewRoleController(RoleUseCase roles.RoleUsecaseInterface) *RoleController{
	return &RoleController{
		roleUseCase: RoleUseCase,
	}
}

func (roleController *RoleController) Add (c echo.Context) error {
	req := request.AddRoleRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data roles.Domain
	data, err = roleController.roleUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainRole(data))

}

func (roleController *RoleController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	role, err := roleController.roleUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAll(role))

}

func (roleController *RoleController) Delete(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = roleController.roleUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}