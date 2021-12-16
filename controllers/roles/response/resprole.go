package response

import (
	"goodjobs/business/roles"
	"time"

	"gorm.io/gorm"
)

type RoleResponse struct {
	Id       	uint         	`json:"id"`
	CreatedAt 	time.Time    	`json:"createdAt"`
	UpdatedAt 	time.Time    	`json:"updatedAt"`
	DeletedAt	gorm.DeletedAt 	`json:"deletedAt"`
	Name     	string       	`json:"name"`

}

func FromDomainRole(domain roles.Domain) RoleResponse {
	return RoleResponse{
		Id 			:domain.Id,
		Name 		:domain.Name,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}
}

func GetAll(domain []roles.Domain) []RoleResponse {
	var GetAll []RoleResponse
	for _, val := range domain{
		GetAll = append(GetAll, FromDomainRole(val))
	}
	return GetAll 
}