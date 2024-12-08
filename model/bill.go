package model

import (
	"time"
)

// BillAggregate is for bill aggregate
type BillAggregate struct {
	Bill
	FeeHistory  *BillFee
	BankAccount *BillBankAccount
}

// BillState is for bill state
type BillState int32

// BillStateConst is for bill state constant
const (
	BillStateNone BillState = iota
	BillStatePending
	BillStateProgress
	BillStatePaid
	BillStateError BillState = 999999
)

var billStateStr []string = []string{"None", "Pending", "Progress", "Paid", "Error"}
var billStateMap = map[string]int{
	"none":     int(BillStateNone),
	"pending":  int(BillStatePending),
	"progress": int(BillStateProgress),
	"paid":     int(BillStatePaid),
	"error":    int(BillStateError),
}

// Bill is model for bill
type Bill struct {
	BaseModel
	ApplicantCompanyID int64     `json:"applicantCompanyID"`
	RecipientCompanyID int64     `json:"recipientCompanyID"`
	IssueDate          time.Time `json:"issueDate"`
	State              BillState `json:"state"`
	PaymentAmount      int64     `json:"paymentAmount"`
	Fee                int64     `json:"fee"`
	FeeRate            int64     `json:"feeRate"`
	FeeRateScale       int       `json:"feeRateScale"`
	Tax                int64     `json:"tax"`
	TaxRate            int64     `json:"taxRate"`
	TaxRateScale       int       `json:"taxRateScale"`
	TotalAmount        int64     `json:"totalAmount"`
	DueDate            time.Time `json:"dueDate"`
}

// NewBill is for new bill
func NewBill(amount int64, dueDate time.Time, fee *Fee) *Bill {
	bill := &Bill{
		IssueDate:     time.Now(),
		State:         BillStatePending,
		PaymentAmount: amount,
		DueDate:       dueDate,
	}

	bill.FeeRate = fee.FeeRate
	bill.FeeRateScale = fee.FeeScale
	bill.TaxRate = fee.TaxRate
	bill.TaxRateScale = fee.TaxScale
	bill.CalculateFee(fee)
	return bill
}

// CalculateFee is for calculate fee
func (b *Bill) CalculateFee(fee *Fee) {
	b.Fee = fee.CalculateFeeAmount(b.PaymentAmount)
	b.Tax = fee.CalculateTaxAmount(b.PaymentAmount)
	b.TotalAmount = b.PaymentAmount + b.Fee + b.Tax
}

// BillFee is for bill fee information
type BillFee struct {
	BillID int64 `json:"billID"`
	Fee
}

// BillBankAccount is for bill account information
type BillBankAccount struct {
	BillID int64 `json:"billID"`
	BankAccount
}
