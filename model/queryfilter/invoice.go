package queryfilter

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// InvoiceQueryFilter is for invoice query filter
type InvoiceQueryFilter struct {
	Filter
	DateFrom  *time.Time `query:"dateFrom"`
	DateTo    *time.Time `query:"dateTo"`
	CompanyID *uint64    `query:"companyID"`
}

// MakeQuery ...
func (f InvoiceQueryFilter) MakeQuery(ctx context.Context, scope interface{}) (interface{}, error) {
	switch scope.(type) {
	case *gorm.DB:
		return f.makeGormQuery(ctx, scope.(*gorm.DB)), nil
	}
	return nil, errors.New("scope is not exist")
}

// makeGormQuery ...
func (f InvoiceQueryFilter) makeGormQuery(_ context.Context, scope *gorm.DB) *gorm.DB {
	if f.DateFrom != nil {
		scope = scope.Where("due_date >= ?", f.DateFrom)
	}
	if f.DateTo != nil {
		scope = scope.Where("due_date <= ?", f.DateTo)
	}
	if f.CompanyID != nil {
		scope = scope.Where("applicant_company_id = ?", f.CompanyID)
	}
	return scope
}
