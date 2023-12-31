// Code generated by mockery v2.28.2. DO NOT EDIT.

package exec

import mock "github.com/stretchr/testify/mock"

// mockReader is an autogenerated mock type for the reader type
type mockReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: p
func (_m *mockReader) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockReader interface {
	mock.TestingT
	Cleanup(func())
}

// newMockReader creates a new instance of mockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockReader(t mockConstructorTestingTnewMockReader) *mockReader {
	mock := &mockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
