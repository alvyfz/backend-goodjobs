package reviews

import (
	"context"
	"errors"
	"goodjobs/business/reviews"

	"gorm.io/gorm"
)

type reviewRepo struct {
	DB *gorm.DB
}

func NewReviewRepo(DB *gorm.DB) *reviewRepo {
	return &reviewRepo{DB: DB}
}

func (Repo *reviewRepo) Add(ctx context.Context, domain reviews.Domain) (reviews.Domain, error) {
	review := Review{
		Id			:domain.Id,
		User_ID		:domain.User_ID,
		Building_ID	:domain.Building_ID,
		Rating		:domain.Rating,
		Description: domain.Description,
	}
	err := Repo.DB.Create(&review)
	if err.Error != nil {
		return reviews.Domain{}, err.Error
	}
	return review.ToDomain(), nil
}

func (Repo *reviewRepo) GetAll(ctx context.Context) ([]reviews.Domain, error){
	var review []Review
	err := Repo.DB.Preload("user").Find(&review)
	if err.Error != nil {
		return []reviews.Domain{}, err.Error
	}
	return GetAll(review), nil
}

func (Repo *reviewRepo) GetByBuildingID(buildingid uint, ctx context.Context ) (reviews.Domain, error){
	var review Review
	err := Repo.DB.Preload("Building").Find(&review, "building_id=?", buildingid)
	if err.Error != nil {
		return reviews.Domain{}, err.Error
	}
	return review.ToDomain(), nil
}

func (Repo *reviewRepo) Edit(id uint, ctx context.Context, domain reviews.Domain) (reviews.Domain, error){
	review := FromDomain(domain)
	if Repo.DB.Save(&review).Error != nil {
		return reviews.Domain{}, errors.New("id tidak ditemukan")
	}
	return review.ToDomain(), nil
}


func (Repo *reviewRepo) Delete(id uint, ctx context.Context) error{
	review := Review{}
	err := Repo.DB.Delete(&review, id)
	if err.Error!= nil {
		return err.Error
		
	}
	if err.RowsAffected == 0 {
		return errors.New("data kosong")
	}
	return nil
}