package mocks

import (
	"context"
	"goodjobs/business/buildings"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) Add(ctx context.Context, domain buildings.Domain) (buildings.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 buildings.Domain
	if rf, ok := ret.Get(0).(func(context.Context, buildings.Domain) buildings.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(buildings.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, buildings.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

func (_m *Repository) GetAll(ctx context.Context) ([]buildings.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []buildings.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []buildings.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buildings.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetOrderByPriceAsc(ctx context.Context) ([]buildings.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []buildings.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []buildings.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buildings.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetOrderByPriceDesc(ctx context.Context) ([]buildings.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []buildings.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []buildings.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buildings.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetByComplexID(complexid uint, ctx context.Context) ([]buildings.Domain, error){
	ret := _m.Called(complexid, ctx)
	var r0 []buildings.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context)[]buildings.Domain); ok{
		r0 = rf(complexid, ctx)
	} else {
		r0 = ret.Get(0).([]buildings.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context)error); ok{
		r1 = rf(complexid, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) GetByID (id uint, ctx context.Context) (buildings.Domain, error){
	ret := _m.Called(id, ctx)
	var r0 buildings.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context)buildings.Domain); ok{
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(buildings.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context)error); ok{
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) Edit(id uint, ctx context.Context, domain buildings.Domain) (buildings.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 buildings.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context, buildings.Domain) buildings.Domain); ok {
		r0 = rf(id, ctx, domain)
	} else {
		r0 = ret.Get(0).(buildings.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context, buildings.Domain) error); ok {
		r1 = rf(id, ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Delete(id uint, ctx context.Context) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}