package request

import "goodjobs/business/users"

type UpdatePasswordReq struct {
	Password 	string 	`json:"password"`
}

func (User *UpdatePasswordReq) ToDomain() *users.Domain {
	return &users.Domain{
		Password:User.Password,
	}
}