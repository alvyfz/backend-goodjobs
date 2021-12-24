package complexes

import (
	"context"
	"errors"
	"time"
)

type ComplexUseCase struct {
	repo ComplexRepoInterface
	ctx  time.Duration
}

func NewUseCase(complexRepo ComplexRepoInterface, ctx time.Duration) *ComplexUseCase {
	return &ComplexUseCase{
		repo: complexRepo,
		ctx:  ctx,
	}
}

func (usecase *ComplexUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("nama complex harus di isi")
	}

	complex, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return complex, nil
}

func (usecase *ComplexUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	complex, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada complex")
	}
	return complex, nil
}

func (usecase *ComplexUseCase) GetByID(id uint, ctx context.Context) (Domain, error){
	complex, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada complex dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return complex, nil
}


func (usecase *ComplexUseCase) Edit(id uint, ctx context.Context, domain Domain) (Domain, error){
	domain.Id = (id)
	complex, err := usecase.repo.Edit(id, ctx , domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada complex dengan ID tersebut")
	}
	return complex, nil
}

func (usecase *ComplexUseCase) Delete(id uint, ctx context.Context) error {
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada Complex dengan ID tersebut")
	}
	return nil
}