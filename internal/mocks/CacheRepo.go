// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	model "github.com/gotway/gotway/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// CacheRepo is an autogenerated mock type for the CacheRepo type
type CacheRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: cache, serviceKey
func (_m *CacheRepo) Create(cache model.Cache, serviceKey string) error {
	ret := _m.Called(cache, serviceKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Cache, string) error); ok {
		r0 = rf(cache, serviceKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByPath provides a mock function with given fields: paths
func (_m *CacheRepo) DeleteByPath(paths []model.CachePath) error {
	ret := _m.Called(paths)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.CachePath) error); ok {
		r0 = rf(paths)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByTags provides a mock function with given fields: tags
func (_m *CacheRepo) DeleteByTags(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: path, serviceKey
func (_m *CacheRepo) Get(path string, serviceKey string) (model.Cache, error) {
	ret := _m.Called(path, serviceKey)

	var r0 model.Cache
	if rf, ok := ret.Get(0).(func(string, string) model.Cache); ok {
		r0 = rf(path, serviceKey)
	} else {
		r0 = ret.Get(0).(model.Cache)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, serviceKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewCacheRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheRepo creates a new instance of CacheRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheRepo(t NewCacheRepoT) *CacheRepo {
	mock := &CacheRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
