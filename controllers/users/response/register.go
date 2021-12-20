package response

import (
	"goodjobs/business/users"
)


func FromUserRegister(domain users.Domain) UserResponse {
	return UserResponse{
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		DeletedAt	:domain.DeletedAt,
		Name		:domain.Name,
		Email		:domain.Email,
		Phone		:domain.Phone,
		Roles_ID	:domain.Roles_ID,
		// Role		:response.FromDomainRole(domain.Roles) ,
	}
}