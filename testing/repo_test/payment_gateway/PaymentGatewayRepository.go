// Code generated by mockery v2.12.2. DO NOT EDIT.

package payment_gateway

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// PaymentGatewayRepository is an autogenerated mock type for the PaymentGatewayRepository type
type PaymentGatewayRepository struct {
	mock.Mock
}

// CreateTransactionRecord provides a mock function with given fields: id, record
func (_m *PaymentGatewayRepository) CreateTransactionRecord(id int, record entities.TransactionRecord) error {
	ret := _m.Called(id, record)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, entities.TransactionRecord) error); ok {
		r0 = rf(id, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIDInvoicePayment provides a mock function with given fields: id
func (_m *PaymentGatewayRepository) GetIDInvoicePayment(id int) (entities.TransactionRecord, error) {
	ret := _m.Called(id)

	var r0 entities.TransactionRecord
	if rf, ok := ret.Get(0).(func(int) entities.TransactionRecord); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.TransactionRecord)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInvoices provides a mock function with given fields: id
func (_m *PaymentGatewayRepository) GetInvoices(id int) (entities.Invoice, []entities.InvoiceItem, error) {
	ret := _m.Called(id)

	var r0 entities.Invoice
	if rf, ok := ret.Get(0).(func(int) entities.Invoice); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entities.Invoice)
	}

	var r1 []entities.InvoiceItem
	if rf, ok := ret.Get(1).(func(int) []entities.InvoiceItem); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]entities.InvoiceItem)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTotalAmount provides a mock function with given fields: id
func (_m *PaymentGatewayRepository) GetTotalAmount(id int) (float32, error) {
	ret := _m.Called(id)

	var r0 float32
	if rf, ok := ret.Get(0).(func(int) float32); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusInvoice provides a mock function with given fields: id, invoice
func (_m *PaymentGatewayRepository) UpdateStatusInvoice(id int, invoice entities.Invoice) error {
	ret := _m.Called(id, invoice)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, entities.Invoice) error); ok {
		r0 = rf(id, invoice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPaymentGatewayRepository creates a new instance of PaymentGatewayRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPaymentGatewayRepository(t testing.TB) *PaymentGatewayRepository {
	mock := &PaymentGatewayRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
