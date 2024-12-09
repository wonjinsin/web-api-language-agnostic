package repository

import (
	"context"
	"pikachu/model"
	"pikachu/model/queryfilter"

	"gorm.io/gorm"
)

type gormInvoiceReadOnlyRepository struct {
	conn *gorm.DB
}

// NewGormInvoiceReadOnlyRepository ...
func NewGormInvoiceReadOnlyRepository(conn *gorm.DB) InvoiceReadOnlyRepository {
	return &gormInvoiceReadOnlyRepository{conn: conn}
}

// GetInvoices ...
func (g *gormInvoiceReadOnlyRepository) GetInvoices(ctx context.Context, filter queryfilter.QueryFilter) (invoices model.InvoiceAggreates, err error) {
	zlog.With(ctx).Infow("[New Repository Request]", "filter", filter)
	scope := g.conn.WithContext(ctx)
	queryScope, err := filter.MakeQuery(ctx, scope)
	if err != nil {
		zlog.With(ctx).Errorw("GetInvoices Error", "filter", filter, "err", err)
		return nil, err
	}

	scope = queryScope.(*gorm.DB)
	scope = scope.
		Preload("Fee").
		Preload("BankAccount").
		Find(&invoices)
	if err = scope.Error; err != nil {
		zlog.With(ctx).Errorw("GetInvoices Error", "filter", filter, "err", err)
		return nil, err
	}

	return invoices, nil
}
