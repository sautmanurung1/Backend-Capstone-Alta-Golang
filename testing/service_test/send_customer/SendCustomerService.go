// Code generated by mockery v2.12.2. DO NOT EDIT.

package send_customer

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	mock "github.com/stretchr/testify/mock"
)

// SendCustomerService is an autogenerated mock type for the SendCustomerService type
type SendCustomerService struct {
	mock.Mock
}

// SendEmailService provides a mock function with given fields: message
func (_m *SendCustomerService) SendEmailService(message entities.SendCustomer) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.SendCustomer) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
