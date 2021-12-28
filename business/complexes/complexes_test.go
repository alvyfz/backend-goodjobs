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

