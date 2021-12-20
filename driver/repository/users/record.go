package users

import (
	"goodjobs/business/users"
	"goodjobs/driver/repository/roles"
	"time"

	"gorm.io/gorm"
)


type User struct {
	Id 			uint 			`gorm:"primaryKey"`
	Email		string			`gorm:"unique"`
	Name 		string 			
	Phone		string
	Password	string
	Roles_ID	uint
	Roles 		roles.Role		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Token		string
	CreatedAt 	time.Time   
	UpdatedAt 	time.Time   
	DeletedAt	gorm.DeletedAt 	`gorm:"index"`
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id 			:user.Id,
		Email		:user.Email,
		Name 		:user.Name,
		Phone		:user.Phone,
		Password	:user.Password,
		Roles_ID	:user.Roles_ID ,
		Roles		:user.Roles.ToDomain() ,
		Token		:user.Token,
		CreatedAt 	:user.CreatedAt,
		UpdatedAt 	:user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User  {
	return User{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name,
		Phone		:domain.Phone,
		Password	:domain.Password,
		Roles_ID	:domain.Roles_ID,
		Roles		:roles.FromDomain(domain.Roles) ,
		Token		:domain.Token,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}

func GetAllUsers(data []User) []users.Domain{
	res := []users.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}