package units

import (
	"context"
	"errors"
	"goodjobs/business/units"

	"gorm.io/gorm"
)

type unitRepo struct {
	DB *gorm.DB
}

func NewUnitRepo(DB *gorm.DB) *unitRepo {
	return &unitRepo{DB: DB}
}

func (Repo *unitRepo) Add(ctx context.Context, domain units.Domain) (units.Domain, error) {
	unit := Unit{
		Id:          domain.Id,
		Name:        domain.Name,
		Building_ID: domain.Building_ID,
		Description: domain.Description,
		Price:       domain.Price,
		UnitSize:    domain.UnitSize,
		Img:         domain.Img,
	}
	err := Repo.DB.Create(&unit)
	if err.Error != nil {
		return units.Domain{}, err.Error
	}
	return unit.ToDomain(), nil
}

func (Repo *unitRepo) GetAll(ctx context.Context) ([]units.Domain, error) {
	var unit []Unit
	err := Repo.DB.Preload("Building").Find(&unit)
	if err.Error != nil {
		return []units.Domain{}, err.Error
	}
	return GetAll(unit), nil
}

func (Repo *unitRepo) Edit(id uint, ctx context.Context, domain units.Domain) (units.Domain, error) {
	unit := FromDomain(domain)
	if Repo.DB.Save(&unit).Error != nil {
		return units.Domain{}, errors.New("id tidak ditemukan")
	}
	return unit.ToDomain(), nil
}

func (Repo *unitRepo) Delete(id uint, ctx context.Context) error {
	unit := Unit{}
	err := Repo.DB.Delete(&unit, id)
	if err.Error != nil {
		return err.Error

	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}