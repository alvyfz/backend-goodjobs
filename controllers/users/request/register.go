package request

import "goodjobs/business/users"

type RegisterUserRequest struct {
	Email    	string 	`json:"email"`
	Name     	string 	`json:"name"`
	Phone		string 	`json:"phone"`
	Password 	string 	`json:"password"`
	Roles_ID	uint	`json:"roles_id"`
}

func (User *RegisterUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email	:User.Email,
		Name    :User.Name,
		Phone	:User.Phone,
		Password:User.Password,
		Roles_ID:User.Roles_ID,
	}
}