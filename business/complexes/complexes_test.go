package complexes_test

import (
	"context"
	"errors"
	"goodjobs/business/complexes"
	"goodjobs/business/complexes/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var complexRepository = mocks.Repository{Mock: mock.Mock{}}
var complexService complexes.ComplexUsecaseInterface
var complexDomain complexes.Domain

var allcomplexDomain []complexes.Domain

func setup() {
	complexService = complexes.NewUseCase(&complexRepository, time.Hour*10)
	complexDomain = complexes.Domain{
		Id:      1,
		Name: "testing",
		Address: "digidaw",
		Img: "foto.com",
	}
}

func TestAdd(t *testing.T) {
	setup()
	complexRepository.On("Add", mock.Anything, mock.Anything).Return(complexDomain, nil)
	complexRepository.On(mock.Anything, mock.Anything).Return(complexes.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		complex, err := complexService.Add(context.Background(), complexes.Domain{
			Id:      1,
			Name: "testing",
			Address: "digidaw",
			Img: "foto.com",
		})

		assert.NoError(t, err)
		assert.Equal(t, complexDomain, complex)
	})

	t.Run("Test Case 2 | No Name", func(t *testing.T) {
		
		complexRepository.On("Add", mock.Anything, mock.Anything).Return(complexDomain, errors.New("nama complex harus di isi")).Once() 
		complex, err := complexService.Add(context.Background(), complexes.Domain{
			Id:      1,
			Name: "",
			Address: "digidaw",
			Img: "foto.com",
		})
		assert.Error(t, err)
		assert.NotNil(t, complex)
	})
	t.Run("Test Case 3 | No Address", func(t *testing.T) {
		
		complexRepository.On("Add", mock.Anything, mock.Anything).Return(complexDomain, errors.New("address harus di isi")).Once() 
		complex, err := complexService.Add(context.Background(), complexes.Domain{
			Id:      1,
			Name: "AWE",
			Address: "",
			Img: "foto.com",
		})
		assert.Error(t, err)
		assert.NotNil(t, complex)
	})
	t.Run("Test Case 4 | No Image", func(t *testing.T) {
		
		complexRepository.On("Add", mock.Anything, mock.Anything).Return(complexDomain, errors.New("image harus di isi")).Once() 
		complex, err := complexService.Add(context.Background(), complexes.Domain{
			Id:      1,
			Name: "AWE",
			Address: "digidaw",
			Img: "",
		})
		assert.Error(t, err)
		assert.NotNil(t, complex)
	})
	
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success Search complex", func(t *testing.T) {
        setup()
        complexRepository.On("GetAll", mock.Anything).Return(allcomplexDomain, nil).Once()
        data, err := complexService.GetAll(context.Background())

        assert.NoError(t, err)
        assert.Nil(t, data)
        assert.Equal(t, len(data), len(allcomplexDomain))
    })

    t.Run("Test case 2 | Error Search complex(search empty)", func(t *testing.T) {
        setup()
        complexRepository.On("GetAll", mock.Anything, mock.Anything).Return([]complexes.Domain{}, errors.New("complex Not Found")).Once()
        data, err := complexService.GetAll(context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []complexes.Domain{})
    })
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetByID", func(t *testing.T) {
		setup()
		complexRepository.On("GetByID", mock.AnythingOfType("uint") ,mock.Anything ).Return(complexDomain, nil).Once()
		complex, err := complexService.GetByID(complexDomain.Id, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, complex)
	})

	t.Run("Test case 2 | Error GetByID(complex Id = 0)", func(t *testing.T) {
		setup()
		complexDomain.Id = 0
		complexRepository.On("GetByID", mock.AnythingOfType("uint") ,mock.Anything ).Return(complexDomain, nil).Once()
		data, err := complexService.GetByID(complexDomain.Id, context.Background())

		assert.Error(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, complexes.Domain{})
	})
}

func TestEdit(t *testing.T) {
	t.Run("Test case 1 | Success Edit", func(t *testing.T) {
		setup()
		complexRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(complexDomain, nil).Once()
		data, err := complexService.Edit(complexDomain.Id, context.Background(), complexDomain )

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Edit", func(t *testing.T) {
		setup()
		complexRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(complexDomain, errors.New("tidak ada complex dengan ID tersebut")).Once()
		data, err := complexService.Edit(complexDomain.Id, context.Background(), complexDomain)

		assert.Equal(t, data, complexes.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		complexRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := complexService.Delete(complexDomain.Id, context.Background() )

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		complexRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("tidak ada Complex dengan ID tersebut")).Once()
		err := complexService.Delete(complexDomain.Id, context.Background())
		assert.Equal(t, err, errors.New("tidak ada Complex dengan ID tersebut"))
		assert.Error(t, err)
	})
}