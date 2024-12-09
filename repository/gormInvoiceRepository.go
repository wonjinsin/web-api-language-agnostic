package repository

import (
	"context"
	"pikachu/model"

	"gorm.io/gorm"
)

type gormInvoiceRepository struct {
	conn *gorm.DB
}

// NewGormInvoiceRepository is for new gorm invoice repository
func NewGormInvoiceRepository(conn *gorm.DB) InvoiceRepository {
	return &gormInvoiceRepository{conn: conn}
}

// NewInvoice is for new invoice
func (g *gormInvoiceRepository) NewInvoice(ctx context.Context, invoice *model.InvoiceAggregate) (rinvoice *model.InvoiceAggregate, err error) {
	zlog.With(ctx).Infow("[New Repository Request]", "invoice", invoice)

	scope := g.conn.WithContext(ctx)
	if err = scope.Create(&invoice).Error; err != nil {
		zlog.With(ctx).Errorw("NewInvoice Error", "err", err)
		return nil, err
	}

	return invoice, nil
}
