package repository

import (
	"context"
	"pikachu/model"
	"pikachu/util"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

// GormFeeReadOnlyRepository ...
type GormFeeReadOnlyRepository struct {
	conn *gorm.DB
}

// NewGormFeeReadOnlyRepository ...
func NewGormFeeReadOnlyRepository(conn *gorm.DB) *GormFeeReadOnlyRepository {
	return &GormFeeReadOnlyRepository{conn: conn}
}

// GetFeeByCountryCode ...
func (g *GormFeeReadOnlyRepository) GetFeeByCountryCode(ctx context.Context, countryCode util.CountryCode) (rfee *model.Fee, err error) {
	zlog.With(ctx).Infow("[New Repository Request]", "countryCode", countryCode)

	scope := g.conn.WithContext(ctx)
	scope = scope.Where("country_code = ?", countryCode).First(&rfee)

	if err = scope.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zlog.With(ctx).Errorw("GetFeeByCountryCode Not Found", "countryCode", countryCode, "err", err)
			return nil, errors.NotFoundf("Fee is not exist")
		}
	}

	return rfee, nil
}
