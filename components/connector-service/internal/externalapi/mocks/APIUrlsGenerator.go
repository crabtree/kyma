// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import httpcontext "github.com/kyma-project/kyma/components/connector-service/internal/clientcontext"
import mock "github.com/stretchr/testify/mock"

// APIUrlsGenerator is an autogenerated mock type for the APIUrlsGenerator type
type APIUrlsGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields: reader
func (_m *APIUrlsGenerator) Generate(reader httpcontext.ConnectorClientReader) interface{} {
	ret := _m.Called(reader)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(httpcontext.ConnectorClientReader) interface{}); ok {
		r0 = rf(reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}
