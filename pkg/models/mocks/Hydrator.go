// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"
import models "github.com/biodebox/go-coge-sql/pkg/models"

// Hydrator is an autogenerated mock type for the Hydrator type
type Hydrator struct {
	mock.Mock
}

// Hydrate provides a mock function with given fields: table, writer
func (_m *Hydrator) Hydrate(table *models.Table, writer io.Writer) error {
	ret := _m.Called(table, writer)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Table, io.Writer) error); ok {
		r0 = rf(table, writer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}