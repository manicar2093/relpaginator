// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package relpaginator

import (
	context "context"

	rel "github.com/go-rel/rel"
	mock "github.com/stretchr/testify/mock"
)

// MockPaginable is an autogenerated mock type for the Paginable type
type MockPaginable struct {
	mock.Mock
}

// CreatePagination provides a mock function with given fields: ctx, tableName, holder, pageSort
func (_m *MockPaginable) CreatePagination(ctx context.Context, tableName string, holder interface{}, pageSort *PageSort) (*Page, error) {
	ret := _m.Called(ctx, tableName, holder, pageSort)

	var r0 *Page
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, *PageSort) *Page); ok {
		r0 = rf(ctx, tableName, holder, pageSort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, *PageSort) error); ok {
		r1 = rf(ctx, tableName, holder, pageSort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePaginator provides a mock function with given fields: ctx, tableName, holder, pageNumber, queries
func (_m *MockPaginable) CreatePaginator(ctx context.Context, tableName string, holder interface{}, pageNumber int, queries ...rel.Querier) (*Page, error) {
	_va := make([]interface{}, len(queries))
	for _i := range queries {
		_va[_i] = queries[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tableName, holder, pageNumber)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *Page
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, int, ...rel.Querier) *Page); ok {
		r0 = rf(ctx, tableName, holder, pageNumber, queries...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, int, ...rel.Querier) error); ok {
		r1 = rf(ctx, tableName, holder, pageNumber, queries...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
