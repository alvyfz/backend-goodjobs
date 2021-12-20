package roles

import (
	"context"
	"errors"
	"time"
)

type RoleUseCase struct {
	repo RoleRepoInterface
	ctx  time.Duration
}

func NewUseCase(roleRepo RoleRepoInterface, ctx time.Duration) *RoleUseCase {
	return &RoleUseCase{
		repo: roleRepo,
		ctx:  ctx,
	}
}

func (usecase *RoleUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("nama role harus di isi")
	}

	role, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return role, nil
}

func (usecase *RoleUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	role, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada buku")
	}
	return role, nil
}

func (usecase *RoleUseCase) Delete(id uint, ctx context.Context) error {
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada Role dengan ID tersebut")
	}
	return nil
}