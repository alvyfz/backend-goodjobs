package units

import (
	"context"
	"errors"
	"time"
)

type UnitUseCase struct {
	repo UnitRepoInterface
	ctx  time.Duration
}

func NewUseCase(unitRepo UnitRepoInterface, ctx time.Duration) *UnitUseCase {
	return &UnitUseCase{
		repo: 		unitRepo,
		ctx:		ctx,
	}
}

func (usecase *UnitUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("nama harus di isi")
	}
	book, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return book, nil
}

func (usecase *UnitUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	book, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada unit")
	}
	return book, nil
}

func (usecase *UnitUseCase) Edit(id uint, ctx context.Context, domain Domain) (Domain, error){
	domain.Id = (id)
	user, err := usecase.repo.Edit(id, ctx , domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada unit dengan ID tersebut")
	}
	return user, nil
}

func (usecase *UnitUseCase) Delete(id uint, ctx context.Context) error{
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada unit dengan ID tersebut")
	}
	return nil
}