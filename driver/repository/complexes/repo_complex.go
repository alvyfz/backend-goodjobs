package complexes

import (
	"context"
	"errors"
	"goodjobs/business/complexes"

	"gorm.io/gorm"
)

type complexRepo struct {
	DB *gorm.DB
}

func NewComplexRepo(DB *gorm.DB) complexes.ComplexRepoInterface {
	return &complexRepo{DB: DB}
}

func (Repo *complexRepo) Add(ctx context.Context, domain complexes.Domain) (complexes.Domain, error) {
	complex := Complex{
		Id 			:domain.Id,
		Name 		:domain.Name,
		Img			:domain.Img,
	}
	err := Repo.DB.Create(&complex)
	if err.Error != nil {
		return complexes.Domain{}, err.Error
	}
	return complex.ToDomain(), nil
}

func (Repo *complexRepo) GetAll(ctx context.Context) ([]complexes.Domain, error){
	var complex []Complex
	err := Repo.DB.Find(&complex)
	if err.Error != nil {
		return []complexes.Domain{}, err.Error
	}
	return GetAll(complex), nil
}

func (Repo *complexRepo) Edit(id uint, ctx context.Context, domain complexes.Domain) (complexes.Domain, error){
	complex := FromDomain(domain)
	if Repo.DB.Save(&complex).Error != nil {
		return complexes.Domain{}, errors.New("id tidak ditemukan")
	}
	return complex.ToDomain(), nil
}


func (Repo *complexRepo) Delete(id uint, ctx context.Context) error{
	complex := Complex{}
	err := Repo.DB.Delete(&complex, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}