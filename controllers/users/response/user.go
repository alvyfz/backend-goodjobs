package response

import (
	"goodjobs/business/users"
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	Id        	uint           `json:"id"`
	CreatedAt 	time.Time      `json:"createdAt"`
	UpdatedAt 	time.Time      `json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt `json:"deletedAt"`
	Email     	string         `json:"email"`
	Name      	string         `json:"name"`
	Phone		string 			`json:"phone"`
	Roles_ID	uint			`json:"role_id"`

}

type JWTResponse struct {
	Token string		`json:"token"`
	User interface{}	`json:"user"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		Name		:domain.Name,
		Email		:domain.Email,
		Phone		:domain.Phone,
		Roles_ID	:domain.Roles_ID,
	}
}

func GetAllUsers(domain []users.Domain) []UserResponse {
	var GetAllUsers []UserResponse
	for _, val := range domain{
		GetAllUsers = append(GetAllUsers, FromDomain(val))
	}
	return GetAllUsers 
}