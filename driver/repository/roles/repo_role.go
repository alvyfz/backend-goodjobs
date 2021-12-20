package roles

import (
	"context"
	"errors"

	"goodjobs/business/roles"

	"gorm.io/gorm"
)

type roleRepo struct {
	DB *gorm.DB
}

func NewRoleRepo(DB *gorm.DB) roles.RoleRepoInterface {
	return &roleRepo{DB: DB}
}

func (Repo *roleRepo) Add(ctx context.Context, domain roles.Domain) (roles.Domain, error) {
	role := Role{
		Id 			:domain.Id,
		Name 		:domain.Name,
	}
	err := Repo.DB.Create(&role)
	if err.Error != nil {
		return roles.Domain{}, err.Error
	}
	return role.ToDomain(), nil
}

func (Repo *roleRepo) GetAll(ctx context.Context) ([]roles.Domain, error){
	var role []Role
	err := Repo.DB.Find(&role)
	if err.Error != nil {
		return []roles.Domain{}, err.Error
	}
	return GetAllRoles(role), nil
}

func (Repo *roleRepo) Delete(id uint, ctx context.Context) error{
	role := Role{}
	err := Repo.DB.Delete(&role, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}