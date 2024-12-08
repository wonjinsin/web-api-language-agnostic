package repository

import (
	"context"
	"pikachu/util"

	"gorm.io/gorm"
)

type gormTxRepository struct {
	conn *gorm.DB
}

// NewGormTxRepository ...
func NewGormTxRepository(conn *gorm.DB) TxRepository {
	return &gormTxRepository{conn: conn}
}

// Begin is for transaction
func (g *gormTxRepository) Begin(ctx context.Context, fn func(ctx context.Context) error) error {
	return g.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		zlog.With(ctx).Infow("[New Repository Request tx]")
		ctx = context.WithValue(ctx, util.TxKey, tx)
		return fn(ctx)
	})
}
