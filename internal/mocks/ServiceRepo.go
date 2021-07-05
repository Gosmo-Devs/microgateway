// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	model "github.com/gotway/gotway/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// ServiceRepo is an autogenerated mock type for the ServiceRepo type
type ServiceRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: service
func (_m *ServiceRepo) Create(service model.Service) error {
	ret := _m.Called(service)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Service) error); ok {
		r0 = rf(service)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: key
func (_m *ServiceRepo) Delete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *ServiceRepo) Get(key string) (model.Service, error) {
	ret := _m.Called(key)

	var r0 model.Service
	if rf, ok := ret.Get(0).(func(string) model.Service); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(model.Service)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *ServiceRepo) GetAll() ([]model.Service, error) {
	ret := _m.Called()

	var r0 []model.Service
	if rf, ok := ret.Get(0).(func() []model.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: service
func (_m *ServiceRepo) Upsert(service model.Service) error {
	ret := _m.Called(service)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Service) error); ok {
		r0 = rf(service)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}