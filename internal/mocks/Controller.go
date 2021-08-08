// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "github.com/gotway/gotway/internal/cache"

	http "net/http"

	mock "github.com/stretchr/testify/mock"

	model "github.com/gotway/gotway/internal/model"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// DeleteCacheByPath provides a mock function with given fields: paths
func (_m *Controller) DeleteCacheByPath(paths []model.CachePath) error {
	ret := _m.Called(paths)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.CachePath) error); ok {
		r0 = rf(paths)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCacheByTags provides a mock function with given fields: tags
func (_m *Controller) DeleteCacheByTags(tags []string) error {
	ret := _m.Called(tags)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCache provides a mock function with given fields: r, service
func (_m *Controller) GetCache(r *http.Request, service string) (model.Cache, error) {
	ret := _m.Called(r, service)

	var r0 model.Cache
	if rf, ok := ret.Get(0).(func(*http.Request, string) model.Cache); ok {
		r0 = rf(r, service)
	} else {
		r0 = ret.Get(0).(model.Cache)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleResponse provides a mock function with given fields: r, params
func (_m *Controller) HandleResponse(r *http.Response, params cache.Params) error {
	ret := _m.Called(r, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Response, cache.Params) error); ok {
		r0 = rf(r, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsCacheableRequest provides a mock function with given fields: r
func (_m *Controller) IsCacheableRequest(r *http.Request) bool {
	ret := _m.Called(r)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*http.Request) bool); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Start provides a mock function with given fields: ctx
func (_m *Controller) Start(ctx context.Context) {
	_m.Called(ctx)
}
