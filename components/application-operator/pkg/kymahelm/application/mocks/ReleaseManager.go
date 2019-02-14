// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import application "github.com/kyma-project/kyma/components/application-operator/pkg/kymahelm/application"
import mock "github.com/stretchr/testify/mock"
import release "k8s.io/helm/pkg/proto/hapi/release"

// ReleaseManager is an autogenerated mock type for the ReleaseManager type
type ReleaseManager struct {
	mock.Mock
}

// CheckReleaseExistence provides a mock function with given fields: name
func (_m *ReleaseManager) CheckReleaseExistence(name string) (bool, error) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckReleaseStatus provides a mock function with given fields: name
func (_m *ReleaseManager) CheckReleaseStatus(name string) (release.Status_Code, string, error) {
	ret := _m.Called(name)

	var r0 release.Status_Code
	if rf, ok := ret.Get(0).(func(string) release.Status_Code); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(release.Status_Code)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteReleaseIfExists provides a mock function with given fields: name
func (_m *ReleaseManager) DeleteReleaseIfExists(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOverridesDefaults provides a mock function with given fields:
func (_m *ReleaseManager) GetOverridesDefaults() application.OverridesData {
	ret := _m.Called()

	var r0 application.OverridesData
	if rf, ok := ret.Get(0).(func() application.OverridesData); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(application.OverridesData)
	}

	return r0
}

// InstallChart provides a mock function with given fields: name, overrides
func (_m *ReleaseManager) InstallChart(name string, overrides string) (release.Status_Code, string, error) {
	ret := _m.Called(name, overrides)

	var r0 release.Status_Code
	if rf, ok := ret.Get(0).(func(string, string) release.Status_Code); ok {
		r0 = rf(name, overrides)
	} else {
		r0 = ret.Get(0).(release.Status_Code)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(name, overrides)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(name, overrides)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
