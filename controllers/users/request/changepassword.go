package request

import "goodjobs/business/users"

type ChangePasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NewPassword string `json:"newpassword"`
}

func ToDomainChangePassword(changepassword ChangePasswordRequest) users.Domain {
	return users.Domain{
		Email:    changepassword.Email,
		Password: changepassword.Password,
		NewPassword: changepassword.NewPassword,
	}
}