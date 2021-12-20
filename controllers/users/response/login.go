package response

import "goodjobs/business/users"

func UserLogin(domain users.Domain, token string) JWTResponse{
	Response := UserResponse{
		Id:        domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Phone: domain.Phone,
		Roles_ID: domain.Roles_ID,
	}

	TokenResponse := JWTResponse{}
	TokenResponse.Token = token
	TokenResponse.User = Response
	return TokenResponse

}

