package buildings

import (
	"context"
	"errors"
	"time"
)

type BuildingUseCase struct {
	repo BuildingRepoInterface
	ctx  time.Duration
}

func NewUseCase(buildingRepo BuildingRepoInterface, ctx time.Duration) *BuildingUseCase {
	return &BuildingUseCase{
		repo: 		buildingRepo,
		ctx:		ctx,
	}
}

func (usecase *BuildingUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("nama harus di isi")
	}
	building, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return building, nil
}

func (usecase *BuildingUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	building, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada building")
	}
	return building, nil
}

func (usecase *BuildingUseCase) GetByID(id uint, ctx context.Context) (Domain, error){
	building, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada building dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return building, nil
}

func (usecase *BuildingUseCase) Edit(id uint, ctx context.Context, domain Domain) (Domain, error){
	domain.Id = (id)
	building, err := usecase.repo.Edit(id, ctx , domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada building dengan ID tersebut")
	}
	return building, nil
}

func (usecase *BuildingUseCase) Delete(id uint, ctx context.Context) error{
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada building dengan ID tersebut")
	}
	return nil
}