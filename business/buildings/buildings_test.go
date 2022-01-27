package buildings_test

import (
	"context"
	"errors"
	"goodjobs/business/buildings"
	"goodjobs/business/buildings/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var buildingRepository = mocks.Repository{Mock: mock.Mock{}}
var buildingService buildings.BuildingUsecaseInterface
var buildingDomain buildings.Domain

var allbuildingDomain []buildings.Domain

func setup() {
	buildingService = buildings.NewUseCase(&buildingRepository, time.Hour*10)
	buildingDomain = buildings.Domain{
		Id:      1,
		Name:    "testing",
		Address: "digidaw",
		Description: "tidakada",
		Size: 12,
		Floor: 5,
		OfficeHours: "libur",
		Toilet: 5,
		Img:     "foto.com",
		Latitude: 5,
		Longitude: 5,
		PriceStart: 5,
	}
}

func TestAdd(t *testing.T) {
	setup()
	buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, nil)
	buildingRepository.On(mock.Anything, mock.Anything).Return(buildings.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})

		assert.NoError(t, err)
		assert.Equal(t, buildingDomain, building)
	})

	t.Run("Test Case 2 | No Name", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("nama harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 3 | No Address", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("alamat harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Name:    "testing",
			Address: "",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 4 | No Image", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("image harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 5 | No Description", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("deskripsi harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 6 | No Complex_ID", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("complex_id harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 0,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 7 | No OfficeHours", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("jam kerja harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 8 | No PriceStart", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("harga harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 0,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 9 | No Toilet", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("toilet harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 0,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 10 | No Floor", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("lantai harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 0,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 11 | No Latitude", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("latitude harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 0,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 12 | No Longitude", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("longitude harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 12,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 0,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})
	t.Run("Test Case 13 | No Size", func(t *testing.T) {

		buildingRepository.On("Add", mock.Anything, mock.Anything).Return(buildingDomain, errors.New("size complex harus di isi")).Once()
		building, err := buildingService.Add(context.Background(), buildings.Domain{
			Id:      1,
			Complex_ID: 1,
			Name:    "testing",
			Address: "digidaw",
			Description: "tidakada",
			Size: 0,
			Floor: 5,
			OfficeHours: "libur",
			Toilet: 5,
			Img:     "foto.com",
			Latitude: 5,
			Longitude: 5,
			PriceStart: 5,
		})
		assert.Error(t, err)
		assert.NotNil(t, building)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success Get building", func(t *testing.T) {
		setup()
		buildingRepository.On("GetAll", mock.Anything).Return(allbuildingDomain, nil).Once()
		data, err := buildingService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Nil(t, data)
		assert.Equal(t, len(data), len(allbuildingDomain))
	})

	t.Run("Test case 2 | Error Get building(get empty)", func(t *testing.T) {
		setup()
		buildingRepository.On("GetAll", mock.Anything, mock.Anything).Return([]buildings.Domain{}, errors.New("building Not Found")).Once()
		data, err := buildingService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []buildings.Domain{})
	})
}

func TestGetOrderByPriceAsc(t *testing.T) {
	t.Run("Test case 1 | Success Get building", func(t *testing.T) {
		setup()
		buildingRepository.On("GetOrderByPriceAsc", mock.Anything).Return(allbuildingDomain, nil).Once()
		data, err := buildingService.GetOrderByPriceAsc(context.Background())

		assert.NoError(t, err)
		assert.Nil(t, data)
		assert.Equal(t, len(data), len(allbuildingDomain))
	})

	t.Run("Test case 2 | Error Get building(get empty)", func(t *testing.T) {
		setup()
		buildingRepository.On("GetOrderByPriceAsc", mock.Anything, mock.Anything).Return([]buildings.Domain{}, errors.New("building Not Found")).Once()
		data, err := buildingService.GetOrderByPriceAsc(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []buildings.Domain{})
	})
}

func TestGetOrderByPriceDesc(t *testing.T) {
	t.Run("Test case 1 | Success Get building", func(t *testing.T) {
		setup()
		buildingRepository.On("GetOrderByPriceDesc", mock.Anything).Return(allbuildingDomain, nil).Once()
		data, err := buildingService.GetOrderByPriceDesc(context.Background())

		assert.NoError(t, err)
		assert.Nil(t, data)
		assert.Equal(t, len(data), len(allbuildingDomain))
	})

	t.Run("Test case 2 | Error Get building(get empty)", func(t *testing.T) {
		setup()
		buildingRepository.On("GetOrderByPriceDesc", mock.Anything, mock.Anything).Return([]buildings.Domain{}, errors.New("building Not Found")).Once()
		data, err := buildingService.GetOrderByPriceDesc(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []buildings.Domain{})
	})
}



func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetByID", func(t *testing.T) {
		setup()
		buildingRepository.On("GetByID", mock.AnythingOfType("uint"), mock.Anything).Return(buildingDomain, nil).Once()
		building, err := buildingService.GetByID(buildingDomain.Id, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, building)
	})

}

func TestEdit(t *testing.T) {
	t.Run("Test case 1 | Success Edit", func(t *testing.T) {
		setup()
		buildingRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(buildingDomain, nil).Once()
		data, err := buildingService.Edit(buildingDomain.Id, context.Background(), buildingDomain)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Edit", func(t *testing.T) {
		setup()
		buildingRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(buildingDomain, errors.New("tidak ada building dengan ID tersebut")).Once()
		data, err := buildingService.Edit(buildingDomain.Id, context.Background(), buildingDomain)

		assert.Equal(t, data, buildings.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		buildingRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := buildingService.Delete(buildingDomain.Id, context.Background())

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		buildingRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("tidak ada building dengan ID tersebut")).Once()
		err := buildingService.Delete(buildingDomain.Id, context.Background())
		assert.Equal(t, err, errors.New("tidak ada building dengan ID tersebut"))
		assert.Error(t, err)
	})
}