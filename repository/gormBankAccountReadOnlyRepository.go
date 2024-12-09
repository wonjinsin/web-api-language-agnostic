package repository

import (
	"context"
	"pikachu/model"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

// GormBankAccountReadOnlyRepository ...
type GormBankAccountReadOnlyRepository struct {
	conn *gorm.DB
}

// NewGormBankAccountReadOnlyRepository ...
func NewGormBankAccountReadOnlyRepository(conn *gorm.DB) *GormBankAccountReadOnlyRepository {
	return &GormBankAccountReadOnlyRepository{conn: conn}
}

// GetBankAccountByCompanyID ...
func (g *GormBankAccountReadOnlyRepository) GetBankAccountByCompanyID(ctx context.Context, companyID uint64) (rbankAccount *model.BankAccount, err error) {
	zlog.With(ctx).Infow("[New Repository Request]", "companyID", companyID)

	scope := g.conn.WithContext(ctx)
	scope = scope.Where("company_id = ?", companyID).First(&rbankAccount)

	if err = scope.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zlog.With(ctx).Errorw("GetBankAccountByCompanyID Not Found", "companyID", companyID, "err", err)
			return nil, errors.NotFoundf("BankAccount is not exist")
		}
	}

	return rbankAccount, nil
}
