// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	accounts "rpcf/accounts"

	mock "github.com/stretchr/testify/mock"
)

// AccountManager is an autogenerated mock type for the AccountManager type
type AccountManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: u
func (_m *AccountManager) Create(u *accounts.Account) (*accounts.Account, error) {
	ret := _m.Called(u)

	var r0 *accounts.Account
	if rf, ok := ret.Get(0).(func(*accounts.Account) *accounts.Account); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accounts.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*accounts.Account) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ID
func (_m *AccountManager) GetByID(ID string) *accounts.Account {
	ret := _m.Called(ID)

	var r0 *accounts.Account
	if rf, ok := ret.Get(0).(func(string) *accounts.Account); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accounts.Account)
		}
	}

	return r0
}

// Login provides a mock function with given fields: u
func (_m *AccountManager) Login(u *accounts.Account) (*accounts.Account, error) {
	ret := _m.Called(u)

	var r0 *accounts.Account
	if rf, ok := ret.Get(0).(func(*accounts.Account) *accounts.Account); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accounts.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*accounts.Account) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
