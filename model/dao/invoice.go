package dao

import "time"

// NewInvoice is for new invoice
type NewInvoice struct {
	ApplicantCompanyID *uint64    `json:"applicantCompanyID"`
	RecipientCompanyID *uint64    `json:"recipientCompanyID"`
	Amount             *int64     `json:"amount"`
	DueDate            *time.Time `json:"dueDate"`
}

// Validate ...
func (n NewInvoice) Validate() bool {
	if n.ApplicantCompanyID == nil || n.RecipientCompanyID == nil || n.Amount == nil || n.DueDate == nil {
		return false
	}
	if *n.ApplicantCompanyID == *n.RecipientCompanyID {
		return false
	}
	if *n.Amount <= 0 || n.DueDate.Before(time.Now()) {
		return false
	}
	return true
}
