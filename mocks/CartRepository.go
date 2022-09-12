// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	cartentity "e-commerce/domains/cart/entity"

	mock "github.com/stretchr/testify/mock"
)

// CartRepoMock is an autogenerated mock type for the IrepoCart type
type CartRepoMock struct {
	mock.Mock
}

// Delete provides a mock function with given fields: cart
func (_m *CartRepoMock) Delete(cart cartentity.CartEntity) error {
	ret := _m.Called(cart)

	var r0 error
	if rf, ok := ret.Get(0).(func(cartentity.CartEntity) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: cart
func (_m *CartRepoMock) Find(cart cartentity.CartEntity) (error, cartentity.CartEntity) {
	ret := _m.Called(cart)

	var r0 error
	if rf, ok := ret.Get(0).(func(cartentity.CartEntity) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	var r1 cartentity.CartEntity
	if rf, ok := ret.Get(1).(func(cartentity.CartEntity) cartentity.CartEntity); ok {
		r1 = rf(cart)
	} else {
		r1 = ret.Get(1).(cartentity.CartEntity)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: cart
func (_m *CartRepoMock) FindAll(cart cartentity.CartEntity) (error, []cartentity.CartEntity) {
	ret := _m.Called(cart)

	var r0 error
	if rf, ok := ret.Get(0).(func(cartentity.CartEntity) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	var r1 []cartentity.CartEntity
	if rf, ok := ret.Get(1).(func(cartentity.CartEntity) []cartentity.CartEntity); ok {
		r1 = rf(cart)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]cartentity.CartEntity)
		}
	}

	return r0, r1
}

// Insert provides a mock function with given fields: cart
func (_m *CartRepoMock) Insert(cart cartentity.CartEntity) error {
	ret := _m.Called(cart)

	var r0 error
	if rf, ok := ret.Get(0).(func(cartentity.CartEntity) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: cart
func (_m *CartRepoMock) Update(cart cartentity.CartEntity) error {
	ret := _m.Called(cart)

	var r0 error
	if rf, ok := ret.Get(0).(func(cartentity.CartEntity) error); ok {
		r0 = rf(cart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCartRepoMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartRepoMock creates a new instance of CartRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartRepoMock(t mockConstructorTestingTNewCartRepoMock) *CartRepoMock {
	mock := &CartRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
