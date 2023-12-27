// Code generated by mockery v1.1.0. DO NOT EDIT.

package labels

import mock "github.com/stretchr/testify/mock"

// MockLabeler is an autogenerated mock type for the Labeler type
type MockLabeler struct {
	mock.Mock
}

// Labels provides a mock function with given fields: otherLabels
func (_m *MockLabeler) Labels(otherLabels map[string]string) map[string]string {
	ret := _m.Called(otherLabels)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(map[string]string) map[string]string); ok {
		r0 = rf(otherLabels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}