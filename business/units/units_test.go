package units_test

import (
	"context"
	"errors"
	"goodjobs/business/units"
	"goodjobs/business/units/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var unitRepository = mocks.Repository{Mock: mock.Mock{}}
var unitService units.UnitUsecaseInterface
var unitDomain units.Domain

var allunitDomain []units.Domain

func setup() {
	unitService = units.NewUseCase(&unitRepository, time.Hour*10)
	unitDomain = units.Domain{
		Id:     1,
		Name:   "testing",
		Building_ID: 1,
		Description: "adanih",
		Price: 1111111,
		UnitSize: 15,
		Img:    "foto.com",
	}
}

func TestAdd(t *testing.T) {
	setup()
	unitRepository.On("Add", mock.Anything, mock.Anything).Return(unitDomain, nil)
	unitRepository.On(mock.Anything, mock.Anything).Return(units.Domain{}, nil)
	t.Run("Test Case 1 | Succss Add", func(t *testing.T) {
		unit, err := unitService.Add(context.Background(), units.Domain{
			Id:     1,
			Name:   "testing",
			Building_ID: 1,
			Description: "adanih",
			Price: 1111111,
			UnitSize: 15,
			Img:    "foto.com",
		})

		assert.NoError(t, err)
		assert.Equal(t, unitDomain, unit)
	})

	t.Run("Test Case 2 | No Name", func(t *testing.T) {

		unitRepository.On("Add", mock.Anything, mock.Anything).Return(unitDomain, errors.New("nama harus di isi")).Once()
		unit, err := unitService.Add(context.Background(), units.Domain{
			Id:     1,
			Name:   "",
			Building_ID: 1,
			Description: "adanih",
			Price: 1111111,
			UnitSize: 15,
			Img:    "foto.com",
		})
		assert.Error(t, err)
		assert.NotNil(t, unit)
	})
	t.Run("Test Case 3 | No Description", func(t *testing.T) {

		unitRepository.On("Add", mock.Anything, mock.Anything).Return(unitDomain, errors.New("deskripsi harus di isi")).Once()
		unit, err := unitService.Add(context.Background(), units.Domain{
			Id:     1,
			Name:   "testing",
			Building_ID: 1,
			Description: "",
			Price: 1111111,
			UnitSize: 15,
			Img:    "foto.com",
		})
		assert.Error(t, err)
		assert.NotNil(t, unit)
	})
	t.Run("Test Case 4 | No Image", func(t *testing.T) {

		unitRepository.On("Add", mock.Anything, mock.Anything).Return(unitDomain, errors.New("img harus di isi")).Once()
		unit, err := unitService.Add(context.Background(), units.Domain{
			Id:     1,
			Name:   "testing",
			Building_ID: 1,
			Description: "adanih",
			Price: 1111111,
			UnitSize: 15,
			Img:    "",
		})
		assert.Error(t, err)
		assert.NotNil(t, unit)
	})
	t.Run("Test Case 5 | No Price", func(t *testing.T) {

		unitRepository.On("Add", mock.Anything, mock.Anything).Return(unitDomain, errors.New("harga harus di isi")).Once()
		unit, err := unitService.Add(context.Background(), units.Domain{
			Id:     1,
			Name:   "testing",
			Building_ID: 1,
			Description: "adanih",
			Price: 0,
			UnitSize: 15,
			Img:    "foto",
		})
		assert.Error(t, err)
		assert.NotNil(t, unit)
	})
	t.Run("Test Case 6 | No Size", func(t *testing.T) {

		unitRepository.On("Add", mock.Anything, mock.Anything).Return(unitDomain, errors.New("size harus di isi")).Once()
		unit, err := unitService.Add(context.Background(), units.Domain{
			Id:     1,
			Name:   "testing",
			Building_ID: 1,
			Description: "adanih",
			Price: 1111111,
			UnitSize: 0,
			Img:    "foto",
		})
		assert.Error(t, err)
		assert.NotNil(t, unit)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Succss Search unit", func(t *testing.T) {
		setup()
		unitRepository.On("GetAll", mock.Anything).Return(allunitDomain, nil).Once()
		data, err := unitService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Nil(t, data)
		assert.Equal(t, len(data), len(allunitDomain))
	})

	t.Run("Test case 2 | Error Search unit(search empty)", func(t *testing.T) {
		setup()
		unitRepository.On("GetAll", mock.Anything, mock.Anything).Return([]units.Domain{}, errors.New("unit Not Found")).Once()
		data, err := unitService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []units.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Succss GetByID", func(t *testing.T) {
		setup()
		unitRepository.On("GetByID", mock.AnythingOfType("uint"), mock.Anything).Return(unitDomain, nil).Once()
		unit, err := unitService.GetByID(unitDomain.Id, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, unit)
	})

	t.Run("Test case 2 | Error GetByID(unit Id = 0)", func(t *testing.T) {
		setup()
		unitDomain.Id = 0
		unitRepository.On("GetByID", mock.AnythingOfType("uint"), mock.Anything).Return(unitDomain, nil).Once()
		data, err := unitService.GetByID(unitDomain.Id, context.Background())

		assert.Error(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, units.Domain{})
	})
}

func TestEdit(t *testing.T) {
	t.Run("Test case 1 | Succss Edit", func(t *testing.T) {
		setup()
		unitRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(unitDomain, nil).Once()
		data, err := unitService.Edit(unitDomain.Id, context.Background(), unitDomain)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Edit", func(t *testing.T) {
		setup()
		unitRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(unitDomain, errors.New("tidak ada unit dengan ID tersebut")).Once()
		data, err := unitService.Edit(unitDomain.Id, context.Background(), unitDomain)

		assert.Equal(t, data, units.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Succss Delete", func(t *testing.T) {
		setup()
		unitRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := unitService.Delete(unitDomain.Id, context.Background())

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		unitRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("tidak ada unit dengan ID tersebut")).Once()
		err := unitService.Delete(unitDomain.Id, context.Background())
		assert.Equal(t, err, errors.New("tidak ada unit dengan ID tersebut"))
		assert.Error(t, err)
	})
}
