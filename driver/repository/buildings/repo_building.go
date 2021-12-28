package buildings

import (
	"context"
	"errors"
	"goodjobs/business/buildings"

	"gorm.io/gorm"
)

type buildingRepo struct {
	DB *gorm.DB
}

func NewBuildingRepo(DB *gorm.DB) *buildingRepo {
	return &buildingRepo{DB: DB}
}

func (Repo *buildingRepo) Add(ctx context.Context, domain buildings.Domain) (buildings.Domain, error) {
	building := Building{
		Complex_ID			:domain.Complex_ID,
		Id					:domain.Id,
		Name				:domain.Name,
		Description			:domain.Description,
		Size				:domain.Size,
		Floor				:domain.Floor,
		OfficeHours			:domain.OfficeHours,
		Address				:domain.Address,
		Toilet				:domain.Toilet,
		Img					:domain.Img,
		Latitude			:domain.Latitude,
		Longitude			:domain.Longitude,
	}
	err := Repo.DB.Create(&building)
	if err.Error != nil {
		return buildings.Domain{}, err.Error
	}
	return building.ToDomain(), nil
}

func (Repo *buildingRepo) GetAll(ctx context.Context) ([]buildings.Domain, error) {
	var building []Building
	err := Repo.DB.Preload("Complex").Find(&building)
	if err.Error != nil {
		return []buildings.Domain{}, err.Error
	}
	return GetAll(building), nil

}

func (Repo *buildingRepo) GetByID(id uint, ctx context.Context ) (buildings.Domain, error){
	var building Building
	err := Repo.DB.Preload("Complex").Find(&building, "id=?", id)
	if err.Error != nil {
		return buildings.Domain{}, err.Error
	}
	return building.ToDomain(), nil
}

func (Repo *buildingRepo) Edit(id uint, ctx context.Context, domain buildings.Domain) (buildings.Domain, error){
	building := FromDomain(domain)
	if Repo.DB.Save(&building).Error != nil {
		return buildings.Domain{}, errors.New("id tidak ditemukan")
	}
	return building.ToDomain(), nil
}


func (Repo *buildingRepo) Delete(id uint, ctx context.Context) error {
	building := Building{}
	err := Repo.DB.Delete(&building, id)
	if err.Error != nil {
		return err.Error

	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}