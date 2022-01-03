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
	if domain.Address == "" {
		return Domain{}, errors.New("alamat harus di isi")
	}
	if domain.Description == "" {
		return Domain{}, errors.New("deskripsi harus di isi")
	}
	if domain.Complex_ID == 0 {
		return Domain{}, errors.New("complex_id harus di isi")
	}
	if domain.OfficeHours == "" {
		return Domain{}, errors.New("jam kerja harus di isi")
	}
	if domain.Img == "" {
		return Domain{}, errors.New("image harus di isi")
	}
	if domain.PriceStart == 0 {
		return Domain{}, errors.New("harga harus di isi")
	}
	if domain.Toilet == 0 {
		return Domain{}, errors.New("toilet harus di isi")
	}
	if domain.Floor == 0 {
		return Domain{}, errors.New("lantai harus di isi")
	}
	if domain.Latitude == 0 {
		return Domain{}, errors.New("latitude harus di isi")
	}
	if domain.Longitude == 0 {
		return Domain{}, errors.New("longitude harus di isi")
	}
	if domain.Size == 0 {
		return Domain{}, errors.New("size complex harus di isi")
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

func (usecase *BuildingUseCase) GetByComplexID(complexid uint, ctx context.Context) ([]Domain, error){
	building, err := usecase.repo.GetByComplexID(complexid, ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada complex dengan ID tersebut")
	}
	if complexid == 0 {
		return []Domain{}, errors.New("complexID harus diisi")
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