// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import pg "github.com/go-pg/pg/v9"

import time "time"

// Listener is an autogenerated mock type for the Listener type
type Listener struct {
	mock.Mock
}

// Channel provides a mock function with given fields:
func (_m *Listener) Channel() <-chan *pg.Notification {
	ret := _m.Called()

	var r0 <-chan *pg.Notification
	if rf, ok := ret.Get(0).(func() <-chan *pg.Notification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *pg.Notification)
		}
	}

	return r0
}

// ChannelSize provides a mock function with given fields: size
func (_m *Listener) ChannelSize(size int) <-chan *pg.Notification {
	ret := _m.Called(size)

	var r0 <-chan *pg.Notification
	if rf, ok := ret.Get(0).(func(int) <-chan *pg.Notification); ok {
		r0 = rf(size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *pg.Notification)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Listener) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Listen provides a mock function with given fields: channels
func (_m *Listener) Listen(channels ...string) error {
	_va := make([]interface{}, len(channels))
	for _i := range channels {
		_va[_i] = channels[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(channels...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Receive provides a mock function with given fields:
func (_m *Listener) Receive() (string, string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReceiveTimeout provides a mock function with given fields: timeout
func (_m *Listener) ReceiveTimeout(timeout time.Duration) (string, string, error) {
	ret := _m.Called(timeout)

	var r0 string
	if rf, ok := ret.Get(0).(func(time.Duration) string); ok {
		r0 = rf(timeout)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(time.Duration) string); ok {
		r1 = rf(timeout)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(time.Duration) error); ok {
		r2 = rf(timeout)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// String provides a mock function with given fields:
func (_m *Listener) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
