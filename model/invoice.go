package model

import (
	"pikachu/util"
	"time"
)

// InvoiceAggregate is for invoice aggregate
type InvoiceAggregate struct {
	Invoice
	Fee         *InvoiceFee         `json:"fee" gorm:"foreignKey:InvoiceID;references:ID"`
	BankAccount *InvoiceBankAccount `json:"bankAccount" gorm:"foreignKey:InvoiceID;references:ID"`
}

// NewInvoiceAggregate is for new invoice aggregate
func NewInvoiceAggregate(
	applicantID string,
	applicantCompanyID uint64,
	recipientCompanyID uint64,
	amount int64,
	dueDate time.Time,
	fee *Fee,
	bankAccount *BankAccount,
) *InvoiceAggregate {
	invoice := newInvoice(
		applicantID,
		applicantCompanyID,
		recipientCompanyID,
		amount,
		dueDate,
		fee,
	)

	newFee := newInvoiceFee(fee)

	newBankAccount := newInvoiceBankAccount(bankAccount)

	return &InvoiceAggregate{
		Invoice:     *invoice,
		Fee:         newFee,
		BankAccount: newBankAccount,
	}
}

// TableName is for table name of gorm
func (InvoiceAggregate) TableName() string {
	return "invoices"
}

// InvoiceAggreates is for invoice aggregates
type InvoiceAggreates []*InvoiceAggregate

// InvoiceState is for invoice state
type InvoiceState int32

// InvoiceStateConst is for invoice state constant
const (
	InvoiceStateNone InvoiceState = iota
	InvoiceStatePending
	InvoiceStateProgress
	InvoiceStatePaid
	InvoiceStateError InvoiceState = 999999
)

var invoiceStateStr []string = []string{"None", "Pending", "Progress", "Paid", "Error"}
var invoiceStateMap = map[string]int{
	"none":     int(InvoiceStateNone),
	"pending":  int(InvoiceStatePending),
	"progress": int(InvoiceStateProgress),
	"paid":     int(InvoiceStatePaid),
	"error":    int(InvoiceStateError),
}

// Invoice is model for invoice
type Invoice struct {
	BaseModel
	ApplicantID        string       `json:"applicantID"`
	ApplicantCompanyID uint64       `json:"applicantCompanyID"`
	RecipientCompanyID uint64       `json:"recipientCompanyID"`
	State              InvoiceState `json:"state"`
	PaymentAmount      int64        `json:"paymentAmount"`
	FeeWithTaxAmount   int64        `json:"feeWithTaxAmount"`
	TotalAmount        int64        `json:"totalAmount"`
	DueDate            time.Time    `json:"dueDate"`
}

func newInvoice(
	applicantID string,
	applicantCompanyID uint64,
	recipientCompanyID uint64,
	amount int64,
	dueDate time.Time,
	fee *Fee,
) *Invoice {
	invoice := &Invoice{
		ApplicantID:        applicantID,
		ApplicantCompanyID: applicantCompanyID,
		RecipientCompanyID: recipientCompanyID,
		State:              InvoiceStatePending,
		PaymentAmount:      amount,
		DueDate:            dueDate,
	}

	invoice.CalculateFee(fee)
	return invoice
}

// CalculateFee is for calculate fee
func (b *Invoice) CalculateFee(fee *Fee) {
	b.FeeWithTaxAmount = fee.CalculateFeeWithTaxAmount(b.PaymentAmount)
	b.TotalAmount = b.PaymentAmount + b.FeeWithTaxAmount
}

// InvoiceFee is for invoice fee information
type InvoiceFee struct {
	ID          uint64           `json:"id"`
	InvoiceID   uint64           `json:"invoiceID"`
	CountryCode util.CountryCode `json:"countryCode"`
	FeeRate     int64            `json:"feeRate"`
	TaxRate     int64            `json:"taxRate"`
	FeeScale    int              `json:"feeScale"`
	TaxScale    int              `json:"taxScale"`
}

func newInvoiceFee(fee *Fee) *InvoiceFee {
	return &InvoiceFee{
		CountryCode: fee.CountryCode,
		FeeRate:     fee.FeeRate,
		TaxRate:     fee.TaxRate,
		FeeScale:    fee.FeeScale,
		TaxScale:    fee.TaxScale,
	}
}

// InvoiceBankAccount is for invoice account information
type InvoiceBankAccount struct {
	ID          uint64 `json:"id"`
	InvoiceID   uint64 `json:"invoiceID"`
	CompanyID   uint64 `json:"companyID"`
	BankName    string `json:"bankName"`
	BranchName  string `json:"branchName"`
	AccountNo   string `json:"accountNo"`
	AccountName string `json:"accountName"`
}

func newInvoiceBankAccount(bankAccount *BankAccount) *InvoiceBankAccount {
	return &InvoiceBankAccount{
		CompanyID:   bankAccount.CompanyID,
		BankName:    bankAccount.BankName,
		BranchName:  bankAccount.BranchName,
		AccountNo:   bankAccount.AccountNo,
		AccountName: bankAccount.AccountName,
	}
}
