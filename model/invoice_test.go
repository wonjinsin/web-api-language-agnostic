package model

import (
	"pikachu/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewInvoiceAggregate(t *testing.T) {
	applicantID := "test-user-id"
	applicantCompanyID := uint64(1)
	recipientCompanyID := uint64(2)
	amount := int64(100000)
	dueDate := time.Now().Add(24 * time.Hour)

	fee := &Fee{
		CountryCode: util.CountryCodeJP,
		FeeRate:     1000,
		TaxRate:     1000,
		FeeScale:    2,
		TaxScale:    2,
	}

	bankAccount := &BankAccount{
		CompanyID:   recipientCompanyID,
		BankName:    "Test Bank",
		BranchName:  "Test Branch",
		AccountNo:   "1234567890",
		AccountName: "Test Account",
	}

	invoice := NewInvoiceAggregate(
		applicantID,
		applicantCompanyID,
		recipientCompanyID,
		amount,
		dueDate,
		fee,
		bankAccount,
	)

	assert.NotNil(t, invoice)
	assert.Equal(t, applicantID, invoice.ApplicantID)
	assert.Equal(t, applicantCompanyID, invoice.ApplicantCompanyID)
	assert.Equal(t, recipientCompanyID, invoice.RecipientCompanyID)
	assert.Equal(t, InvoiceStatePending, invoice.State)
	assert.Equal(t, amount, invoice.PaymentAmount)
	assert.Equal(t, dueDate.Unix(), invoice.DueDate.Unix())

	assert.NotNil(t, invoice.Fee)
	assert.Equal(t, util.CountryCodeJP, invoice.Fee.CountryCode)
	assert.Equal(t, int64(1000), invoice.Fee.FeeRate)
	assert.Equal(t, int64(1000), invoice.Fee.TaxRate)

	assert.NotNil(t, invoice.BankAccount)
	assert.Equal(t, recipientCompanyID, invoice.BankAccount.CompanyID)
	assert.Equal(t, "Test Bank", invoice.BankAccount.BankName)
	assert.Equal(t, "Test Branch", invoice.BankAccount.BranchName)
	assert.Equal(t, "1234567890", invoice.BankAccount.AccountNo)
	assert.Equal(t, "Test Account", invoice.BankAccount.AccountName)
}

func TestInvoiceState_String(t *testing.T) {
	tests := []struct {
		name     string
		state    InvoiceState
		expected string
	}{
		{"State None", InvoiceStateNone, "None"},
		{"State Pending", InvoiceStatePending, "Pending"},
		{"State Progress", InvoiceStateProgress, "Progress"},
		{"State Paid", InvoiceStatePaid, "Paid"},
		{"State Error", InvoiceStateError, "Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.state.String())
		})
	}
}

func TestInvoice_CalculateFee(t *testing.T) {
	invoice := &Invoice{
		PaymentAmount: 100000,
	}

	fee := &Fee{
		CountryCode: util.CountryCodeJP,
		FeeRate:     1000,
		TaxRate:     1000,
		FeeScale:    2,
		TaxScale:    2,
	}

	invoice.CalculateFee(fee)

	expectedFeeWithTax := int64(11000)
	expectedTotal := invoice.PaymentAmount + expectedFeeWithTax

	assert.Equal(t, expectedFeeWithTax, invoice.FeeWithTaxAmount)
	assert.Equal(t, expectedTotal, invoice.TotalAmount)
}
