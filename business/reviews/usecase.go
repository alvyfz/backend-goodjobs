package reviews

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type ReviewUseCase struct {
	repo ReviewRepoInterface
	ctx  time.Duration
}

func NewUseCase(reviewRepo ReviewRepoInterface, ctx time.Duration) *ReviewUseCase {
	return &ReviewUseCase{
		repo: 		reviewRepo,
		ctx:		ctx,
	}
}


func (usecase *ReviewUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	fmt.Println("usecase", domain)
	if domain.Rating == 0 {
		return Domain{}, errors.New("rating harus di isi")
	}
	review, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return review, nil
}

func (usecase *ReviewUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	review, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada review")
	}
	return review, nil
}

func (usecase *ReviewUseCase) GetByBuildingID(buildingid uint, ctx context.Context) (Domain, error){
	unit, err := usecase.repo.GetByBuildingID(buildingid, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada building dengan ID tersebut")
	}
	if buildingid == 0 {
		return Domain{}, errors.New("buildingID harus diisi")
	}
	return unit, nil
}

func (usecase *ReviewUseCase) Edit(id uint, ctx context.Context, domain Domain) (Domain, error){
	domain.Id = (id)
	review, err := usecase.repo.Edit(id, ctx , domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada review dengan ID tersebut")
	}
	return review, nil
}

func (usecase *ReviewUseCase) Delete(id uint, ctx context.Context) error{
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada review dengan ID tersebut")
	}
	return nil
}