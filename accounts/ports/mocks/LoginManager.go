// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	accounts "rpcf/accounts"

	mock "github.com/stretchr/testify/mock"
)

// LoginManager is an autogenerated mock type for the LoginManager type
type LoginManager struct {
	mock.Mock
}

// Login provides a mock function with given fields: acc
func (_m *LoginManager) Login(acc *accounts.Account) (*accounts.Account, string, error) {
	ret := _m.Called(acc)

	var r0 *accounts.Account
	if rf, ok := ret.Get(0).(func(*accounts.Account) *accounts.Account); ok {
		r0 = rf(acc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accounts.Account)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*accounts.Account) string); ok {
		r1 = rf(acc)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*accounts.Account) error); ok {
		r2 = rf(acc)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
