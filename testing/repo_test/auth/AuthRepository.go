// Code generated by mockery v2.12.2. DO NOT EDIT.

package auth

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// GetUserRepository provides a mock function with given fields: id_pegawai
func (_m *AuthRepository) GetUserRepository(id_pegawai int) (entities.Admin, error) {
	ret := _m.Called(id_pegawai)

	var r0 entities.Admin
	if rf, ok := ret.Get(0).(func(int) entities.Admin); ok {
		r0 = rf(id_pegawai)
	} else {
		r0 = ret.Get(0).(entities.Admin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_pegawai)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginRepository provides a mock function with given fields: id_pegawai
func (_m *AuthRepository) LoginRepository(id_pegawai int) (entities.Admin, error) {
	ret := _m.Called(id_pegawai)

	var r0 entities.Admin
	if rf, ok := ret.Get(0).(func(int) entities.Admin); ok {
		r0 = rf(id_pegawai)
	} else {
		r0 = ret.Get(0).(entities.Admin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_pegawai)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterRepository provides a mock function with given fields: _a0
func (_m *AuthRepository) RegisterRepository(_a0 entities.Admin) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.Admin) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
