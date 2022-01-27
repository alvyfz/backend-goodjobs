package reviews_test

import (
	"context"
	"errors"
	"goodjobs/business/reviews"
	"goodjobs/business/reviews/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reviewRepository = mocks.Repository{Mock: mock.Mock{}}
var reviewService reviews.ReviewUsecaseInterface
var reviewDomain reviews.Domain

var allreviewDomain []reviews.Domain

func setup() {
	reviewService = reviews.NewUseCase(&reviewRepository, time.Hour*10)
	reviewDomain = reviews.Domain{
		Id:      1,
		Rating: 50,
		Description: "ada",
		User_ID: 0,
		Building_ID: 0,
	}
}

func TestAdd(t *testing.T) {
	setup()
	reviewRepository.On("Add", mock.Anything, mock.Anything).Return(reviewDomain, nil)
	reviewRepository.On(mock.Anything, mock.Anything).Return(reviews.Domain{}, nil)
	t.Run("Test Case 1 | Success Add", func(t *testing.T) {
		review, err := reviewService.Add(context.Background(), reviews.Domain{
			Id:      1,
			Rating: 50,
			Description: "ada",
			User_ID: 1,
			Building_ID: 1,
		})

		assert.NoError(t, err)
		assert.Equal(t, reviewDomain, review)
	})

	t.Run("Test Case 2 | No Rating", func(t *testing.T) {
		
		reviewRepository.On("Add", mock.Anything, mock.Anything).Return(reviewDomain, errors.New("rating harus di isi")).Once() 
		review, err := reviewService.Add(context.Background(), reviews.Domain{
			Id:      1,
			Rating: 0,
			Description: "ada",
			User_ID: 1,
			Building_ID: 1,
		})
		assert.Error(t, err)
		assert.NotNil(t, review)
	})
	t.Run("Test Case 3 | No Description", func(t *testing.T) {
		
		reviewRepository.On("Add", mock.Anything, mock.Anything).Return(reviewDomain, errors.New("deskripsi harus di isi")).Once() 
		review, err := reviewService.Add(context.Background(), reviews.Domain{
			Id:      1,
			Rating: 50,
			Description: "",
			User_ID: 1,
			Building_ID: 1,
		})
		assert.Error(t, err)
		assert.NotNil(t, review)
	})
	t.Run("Test Case 4 | No UserID", func(t *testing.T) {
		
		reviewRepository.On("Add", mock.Anything, mock.Anything).Return(reviewDomain, errors.New("userid harus di isi")).Once() 
		review, err := reviewService.Add(context.Background(), reviews.Domain{
			Id:      1,
			Rating: 50,
			Description: "ada",
			User_ID: 0,
			Building_ID: 1,
		})
		assert.Error(t, err)
		assert.NotNil(t, review)
	})
	t.Run("Test Case 5 | No BuildingID", func(t *testing.T) {
		
		reviewRepository.On("Add", mock.Anything, mock.Anything).Return(reviewDomain, errors.New("buildingid harus di isi")).Once() 
		review, err := reviewService.Add(context.Background(), reviews.Domain{
			Id:      1,
			Rating: 50,
			Description: "ada",
			User_ID: 1,
			Building_ID: 0,
		})
		assert.Error(t, err)
		assert.NotNil(t, review)
	})
	
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Success Search review", func(t *testing.T) {
        setup()
        reviewRepository.On("GetAll", mock.Anything).Return(allreviewDomain, nil).Once()
        data, err := reviewService.GetAll(context.Background())

        assert.NoError(t, err)
        assert.Nil(t, data)
        assert.Equal(t, len(data), len(allreviewDomain))
    })

    t.Run("Test case 2 | Error Search review(search empty)", func(t *testing.T) {
        setup()
        reviewRepository.On("GetAll", mock.Anything, mock.Anything).Return([]reviews.Domain{}, errors.New("review Not Found")).Once()
        data, err := reviewService.GetAll(context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []reviews.Domain{})
    })
}

func TestGetByBuildingId(t *testing.T) {
	t.Run("Test case 1 | Success GetByID", func(t *testing.T) {
		setup()
		reviewRepository.On("GetByBuildingID", mock.AnythingOfType("uint"), mock.Anything).Return([]reviews.Domain{reviewDomain}, nil).Once()
		review, err := reviewService.GetByBuildingID(reviewDomain.Id, context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, review)
		assert.Contains(t, review, reviewDomain)
	})

	t.Run("Test case 2 | Error Search review(search empty)", func(t *testing.T) {
        setup()
        reviewRepository.On("GetByBuildingID", mock.Anything, mock.Anything).Return([]reviews.Domain{}, errors.New("review Not Found")).Once()
        data, err := reviewService.GetByBuildingID(reviewDomain.Id, context.Background())

        assert.Error(t, err)
        assert.Equal(t, data, []reviews.Domain{})
    })

}

func TestEdit(t *testing.T) {
	t.Run("Test case 1 | Success Edit", func(t *testing.T) {
		setup()
		reviewRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(reviewDomain, nil).Once()
		data, err := reviewService.Edit(reviewDomain.Id, context.Background(), reviewDomain )

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Edit", func(t *testing.T) {
		setup()
		reviewRepository.On("Edit", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(reviewDomain, errors.New("tidak ada review dengan ID tersebut")).Once()
		data, err := reviewService.Edit(reviewDomain.Id, context.Background(), reviewDomain)

		assert.Equal(t, data, reviews.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		reviewRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := reviewService.Delete(reviewDomain.Id, context.Background() )

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		reviewRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("tidak ada review dengan ID tersebut")).Once()
		err := reviewService.Delete(reviewDomain.Id, context.Background())
		assert.Equal(t, err, errors.New("tidak ada review dengan ID tersebut"))
		assert.Error(t, err)
	})
}